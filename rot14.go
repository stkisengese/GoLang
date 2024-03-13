package piscine

func Rot14(s string) string {
        var result string
        for _, r := range s {
                switch {
                case r >= 'a' && r <= 'z':
                        result += string((r-'a'+14)%26 + 'a')
                case r >= 'A' && r <= 'Z': 
                        result += string((r-'A'+14)%26 + 'A') 
                default:
                        result += string(r)
                }
        }
        return result
}