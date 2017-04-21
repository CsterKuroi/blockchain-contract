package pipelines

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"os"

	"unicontract/src/core/model"
	"unicontract/src/common"
	"unicontract/src/config"
	r "unicontract/src/core/db/rethinkdb"
)

func cvChangefeed(in io.Reader, out io.Writer) {
	var value interface{}
	res := r.Changefeed("Unicontract", "Contracts")
	for res.Next(&value) {
		m := value.(map[string]interface{})
		v, err := json.Marshal(m["new_val"])
		if err != nil {
			log.Fatalf(err.Error())
			continue
		}
		if bytes.Equal(v, []byte("null")) {
			continue
		}
		out.Write(v)
	}
}

func cvValidateContract(in io.Reader, out io.Writer) {
	rd := bufio.NewReader(in)
	p := make([]byte, MaxSizeTX)
	for {
		n, _ := rd.Read(p)
		if n == 0 {
			continue
		}
		t := p[:n]
		mod := model.ContractModel{}
		err := json.Unmarshal(t,&mod)
		if err != nil {
			log.Fatalf(err.Error())
			continue
		}
		v := model.Vote{}
		//TODO Validate contract
		if mod.Validate() {
			//vote true
			v.VoteBody.IsValid = true
		} else {
			//vote flase
			v.VoteBody.IsValid = false
		}
		v.VoteBody.VoteForContract = mod.Id
		out.Write([]byte(v.ToString()))
	}
}

func cvVote(in io.Reader, out io.Writer) {
	rd := bufio.NewReader(in)
	p := make([]byte, MaxSizeTX)
	for {
		n, _ := rd.Read(p)
		if n == 0 {
			continue
		}
		t := p[:n]
		v :=model.Vote{}
		err := json.Unmarshal(t,&v)
		if err != nil {
			log.Fatalf(err.Error())
			continue
		}
		//TODO
		v.NodePubkey = config.Config.Keypair.Public
		v.Id = v.GenerateId()
		v.Signature = v.SignVote()
		v.VoteBody.Timestamp = common.GenTimestamp()
		v.VoteBody.VoteType = "node"
		out.Write([]byte(v.ToString()))
	}
}

func cvWriteVote(in io.Reader, out io.Writer) {
	rd := bufio.NewReader(in)
	p := make([]byte, MaxSizeTX)
	for {
		n, _ := rd.Read(p)
		if n == 0 {
			continue
		}
		t := p[:n]
		r.Insert("Unicontract", "Votes", string(t))
		out.Write(t)
	}
}

func startContractVote() {

	p := Pipe(
		cvChangefeed,
		cvValidateContract,
		cvVote,
		cvWriteVote)

	f, err := os.OpenFile("/dev/null", os.O_RDWR, 0)
	if err != nil {
		log.Fatalf(err.Error())
	}
	w := bufio.NewWriter(f)
	p(nil, w)
}
