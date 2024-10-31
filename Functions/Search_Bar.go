package functions

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

var SearchedArtist []Artist
func Search_Bar(w http.ResponseWriter,r *http.Request){
	SearchedArtist:=[]Artist{}
	if r.Method!="GET"{
		ErrorHandler(w,r,405)
		return
	}
	er,err:=Fitch_Global(w,r,Url_Artists),Fitch_Global(w,r,Url_Locations)
	if er!=nil || err!=nil{
		ErrorHandler(w,r,500)
		return
	}
	Search := r.FormValue("search")
	if Search == "" {
		ErrorHandler(w,r,400)
		return
	}
	founded := map[int]bool{}
	for _,artist:=range Artists{
		if strings.Contains(strings.ToLower(artist.Name),strings.ToLower(Search)) && !founded[artist.Id]{
			SearchedArtist=append(SearchedArtist,artist)
			founded[artist.Id]=true
		}
		if strings.Contains(strings.ToLower(artist.FirstAlbum),strings.ToLower(Search)) && !founded[artist.Id]{
			SearchedArtist=append(SearchedArtist,artist)
			founded[artist.Id]=true
		}
		if strings.Contains(strings.ToLower(strconv.Itoa(artist.CreationDate)),strings.ToLower(Search)) && !founded[artist.Id]{
			SearchedArtist=append(SearchedArtist,artist)
			founded[artist.Id]=true
		}
		for _, member := range artist.Members{
			if strings.Contains(strings.ToLower(member),strings.ToLower(Search)) && !founded[artist.Id]{
				SearchedArtist=append(SearchedArtist,artist)
				founded[artist.Id]=true
			}
		}
		for _, location := range Locations.Index {
			for _, loc := range location.Locatins {
				if strings.Contains(strings.ToLower(loc), strings.ToLower(Search)) && !founded[artist.Id] {
					if artist.Id == location.Id  {
					SearchedArtist = append(SearchedArtist, artist)
					founded[artist.Id] = true
					}
				}
			}
		}
	}
	

	tmp, err := template.ParseFiles("Template/Search_Bar.html")
	if err != nil {
		fmt.Println(err)
		fmt.Println("hh")
		ErrorHandler(w,r,500)
		return
	}
	err = tmp.Execute(w, SearchedArtist)
	if err != nil {
		fmt.Println(err)
		fmt.Println("ff")
		ErrorHandler(w,r,500)
		return
	}
}