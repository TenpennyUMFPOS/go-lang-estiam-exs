## BOUTHOURI Mohamed (Aziz) 
## HOW TO RUN
     - go run .
     - ADD : curl "http://localhost:8090/add?nom=peter&definition=parker%20spiderman"
     - REMOVE : curl "http://localhost:8090/remove?nom=formateur"
     - GET : curl "http://localhost:8090/get?nom=peter"
     - UPATE : curl "http://localhost:8090/update?nom=example&newDefinition=newValue"


## HELP SOURCES 
    - https://transform.tools/json-to-go
    - https://pkg.go.dev/net/http
    - https://stackoverflow.com/questions/25465566/golang-parse-json-array-into-data-structure
    - https://medium.com/goturkiye/concurrency-in-go-channels-and-waitgroups-25dd43064d1
    - https://gobyexample.com/waitgroups

## Utilisation du Go routine :
    - mise à jour du dictionnaire :
        go func() {
            for entry := range addChannel {
                dictionary.Add(entry.Nom, entry.Definition)
            }
        }()

    - creation enteriesToAdd pour la parcourir : 
        	entriesToAdd := []packdict.DictionaryEntry{
		        {Nom: "exemple", Definition: "exemple def"},
		        {Nom: "platform", Definition: "teams"},
		        {Nom: "formateur", Definition: "Aziz"},
	        }
    - utilisation de enteriesToAdd : 
            for _, entry := range entriesToAdd {
		        wg.Add(1)
		        go func(e packdict.DictionaryEntry) {
			        defer wg.Done()
			        fmt.Println("chanel working ..")
			        fmt.Printf("Sending entry: %s\n", e.Nom)
			        addChannel <- e
		    }(entry)
	}

	wg.Wait()

	close(addChannel)  