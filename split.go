package piscine

func Split(s, sep string) []string {
    result := [] string{}
    subStr := 0
    for i := 0; i < len(s); i++ {
        if i + len(sep) <= len(s) && s[i:i+len(sep)] == sep {
             result = append(result, s[subStr:i])
        subStr = i+len(sep)
          i += len(sep) - 1         
        }
       } 
          if subStr < len(s) {
            result = append(result, s[subStr:])
          }
    return result
}
