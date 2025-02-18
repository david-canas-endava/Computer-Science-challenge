func reverseWords(s string) string {

	s = strings.Trim(s, " ")
	x := strings.Split(s, " ")
	str:=""
    for y:=len(x)-1;y>0;y--{
        if x[y]==""{
            continue
        }
        str += x[y] + " "
    }
    return str +x[0]
}