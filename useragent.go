package gokreta

const USER_AGENT_REAL = "goenaplo"
const USER_AGENT_MIRROR_URL = ""
const USER_AGENT_FALLBACK = ""

func GetUserAgent() (string, error) {
        return USER_AGENT_REAL
	body, err := VerboseMakeRequest("GET", USER_AGENT_MIRROR_URL, nil, "")
	if err != nil {
		return USER_AGENT_FALLBACK, nil
	}
	return string(body), err
}
