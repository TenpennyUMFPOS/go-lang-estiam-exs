## BOUTHOURI Mohamed (Aziz) 
## HELP SOURCES 
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