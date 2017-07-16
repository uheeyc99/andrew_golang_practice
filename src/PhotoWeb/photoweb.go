package main
import (

	"log"
	"os"
	"net/http"
	"fmt"
	"io/ioutil"
	"html/template"
	"io"
	"path"
	"runtime/debug"
)
//http://blog.csdn.net/tank_war/article/details/54971809
const (
	UPLOAD_DIR = "uploads"
	ListDir = 0x0001
	TEMPLATE_DIR = "./html"
	DEBUG_HTML = true
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("uploadHandler")

	if r.Method == "GET" {

		if err:=renderHtml(w,"upload.html",nil); err !=nil{
			http.Error(w,err.Error(),http.StatusInternalServerError)
		}
		return
	}

	if r.Method == "POST" {
		file1, fileheader, err := r.FormFile("image")
		if err != nil {
			http.Error(w, err.Error(),
				http.StatusInternalServerError)
			return
		}
		filename := fileheader.Filename
		defer file1.Close()

		t, err := os.Create(UPLOAD_DIR + "/" + filename)
		if err != nil {
			log.Println("create dir has error");

			http.Error(w, err.Error(),
				http.StatusInternalServerError)
			return
		}
		defer t.Close()
		if _, err := io.Copy(t, file1); err != nil {
			http.Error(w, err.Error(),
				http.StatusInternalServerError)
			return
		}
		//https://zhidao.baidu.com/question/1574962284977388260.html
		http.Redirect(w, r, "/view5?id="+filename, http.StatusFound)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("viewHandler")
	imageId := r.FormValue("id")
	imagePath := UPLOAD_DIR + "/" + imageId
	if exists := isExists(imagePath);!exists {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, imagePath)

}

func isExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}


func listHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("listHandler")
	fileInfoArr, err := ioutil.ReadDir("uploads")
	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
		return
	}

	locals := make(map[string]interface{})
	images := []string{}
	for _, fileInfo := range fileInfoArr {
		images = append(images, fileInfo.Name())
	}

	locals["images"] = images

	if err=renderHtml(w,"list.html",locals); err !=nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
	}



}

func renderHtml(w http.ResponseWriter,tmpl string, locals map[string]interface{})  error{

	if DEBUG_HTML == true {
		templatePath:= TEMPLATE_DIR + "/" + tmpl
		t,err:=template.ParseFiles(templatePath)
		err =t.Execute(w,locals)
		return err
	}else{
		err := templates[tmpl].Execute(w,locals)
		return  err
	}

}

var templates = make(map[string] *template.Template)



func  prepareHtml(){
	fileInfoArr, err := ioutil.ReadDir(TEMPLATE_DIR)
	if err!=nil{
		panic(err)
		return
	}
	var templateName, templatePath string
	for _,fileInfo :=range fileInfoArr{
		templateName = fileInfo.Name()
		if ext:=path.Ext(templateName);ext !=".html" {
			continue
		}
		templatePath = TEMPLATE_DIR + "/" + templateName
		//template.Must()确保了模板不能解析成功时，一定会触发错误处理流程
		t:=template.Must(template.ParseFiles(templatePath))
		templates[templateName] = t
		log.Println("Loading template:", templatePath,templateName)



	}
	log.Println(templates)

}


func init()  {
	prepareHtml()

}

func safeHandler(fn http.HandlerFunc) http.HandlerFunc{
	return  func(w http.ResponseWriter, r *http.Request){
		defer func() {
			if e,ok:=recover().(error);ok{
				http.Error(w,e.Error(),http.StatusInternalServerError)
				// w.WriteHeader(http.StatusInternalServerError)
				// renderHtml(w, "error", e)
				log.Println("WARN: panic in %v - %v", fn, e)

				log.Println(string(debug.Stack()))
			}

		}()
		fmt.Println(fn)
		fn(w,r)
	}

}


func main() {

	http.HandleFunc("/", safeHandler(listHandler))  //用闭包
	http.HandleFunc("/upload", safeHandler(uploadHandler))
	http.HandleFunc("/view5", safeHandler(viewHandler))
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	//go http.ListenAndServe(":9091",http.FileServer(http.Dir("./")))

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}





}