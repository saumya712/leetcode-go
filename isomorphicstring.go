func isIsomorphic(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    mapST := make(map[byte]byte)
    mapTS := make(map[byte]byte)

    for i := 0; i < len(s); i++ {
        c1 := s[i]
        c2 := t[i]

        // check s -> t mapping
        if val, ok := mapST[c1]; ok {
            if val != c2 {
                return false
            }
        } else {
            mapST[c1] = c2
        }

        // check t -> s mapping
        if val, ok := mapTS[c2]; ok {
            if val != c1 {
                return false
            }
        } else {
            mapTS[c2] = c1
        }
    }

    return true
}
