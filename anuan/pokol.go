package anuan

var akumohon string

func init(){
	akumohon = "Kehidupan Bernegara"
}

func Hailo(Name string) string{
	return "Hellow From Anuan, " + Name + akumohon
}