
type Solution struct{}

func (s *Solution) Encode(strs []string) string {
    var builder strings.Builder
    for _, str := range strs {
        builder.WriteString(str)
        builder.WriteString("#!@")
    }
    return builder.String()
}

func (s *Solution) Decode(encoded string) []string {
    if encoded == "" {
        return []string{}
    }
    res := strings.Split(encoded, "#!@")
    return res[:len(res) - 1]
}
