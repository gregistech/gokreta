package gokreta

const USER_AGENT_MIRROR_URL = "https://raw.githubusercontent.com/boapps/kreta-api-mirror/master/user-agent"
const USER_AGENT_FALLBACK = "Kreta.Ellenorzo/2.9.4.2019101401 (Android; G8341 0.0)"

func GetUserAgent() (string, error) {
	body, err := VerboseMakeRequest("GET", USER_AGENT_MIRROR_URL, nil, "")
	if err != nil {
		return USER_AGENT_FALLBACK, nil
	}
	return string(body), err
}
