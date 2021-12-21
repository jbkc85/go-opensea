package opensea

/**
   "creator":{
     "user":{
        "username":"AlgoTwo"
     },
     "profile_img_url":"https://storage.googleapis.com/opensea-static/opensea-profile/20.png",
     "address":"0x00000444e5a1a667663b0adfd853e8efa0470698",
     "config":""
  },
*/

type Creator struct {
	Address       string            `json:"address"`
	Config        string            `json:"config"`
	ProfileImgUrl string            `json:"profile_img_url"`
	User          map[string]string `json:"user"`
}
