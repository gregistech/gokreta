package gokreta

const USER_AGENT_REAL = "gokreta"
const USER_AGENT_MIRROR_URL = ""
const USER_AGENT_FALLBACK = "gokreta"

// GetUserAgent will return a user agent string we can use
// to access the API.
// Lately there have been blocking by UserAgent and some controversies,
// but right now the situation is calm, so we just use "gokreta".
func GetUserAgent() (string, error) {
	//TODO: Specifiable UserAgent!
	return USER_AGENT_REAL, nil
	body, err := VerboseMakeRequest("GET", USER_AGENT_MIRROR_URL, nil, "")
	if err != nil {
		return USER_AGENT_FALLBACK, err
	}
	return string(body), err
}
