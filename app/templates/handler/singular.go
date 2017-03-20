package handler

import (
    "io"
    "io/ioutil"
	"net/http"
    "encoding/json"

    "github.com/gorilla/mux"
    "<%= vcs %>/<%= repo %>/<%= project %>/utils"
    "<%= vcs %>/<%= repo %>/<%= project %>/models"
    "<%= vcs %>/<%= repo %>/<%= project %>/database"
)

func Retrieve<%= nounSingularUpper %>(db *database.Database, w http.ResponseWriter, r *http.Request) error {
    id := mux.Vars(r)["id"]
    nounSingularLower, err := db.RetrieveOne(id)
    if err != nil {
        return utils.StatusError{http.StatusInternalServerError, err}
    }
    w.Header().Set("Content-Type", "application/json")
    if nounSingularLower != nil {
        w.WriteHeader(http.StatusOK)
        if err := json.NewEncoder(w).Encode(nounSingularLower); err != nil {
            return utils.StatusError{http.StatusInternalServerError, err}
        }
    } else {
        w.WriteHeader(http.StatusNotFound)
    }
    return nil
}

func Update<%= nounSingularUpper %>(db *database.Database, w http.ResponseWriter, r *http.Request) error {
    id := mux.Vars(r)["id"]
    var nounSingularLower models.<%= nounSingularUpper %>
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        return utils.StatusError{http.StatusInternalServerError, err}
    }
	if err := r.Body.Close(); err != nil {
        return utils.StatusError{http.StatusInternalServerError, err}
	}
	if err := json.Unmarshal(body, &nounSingularLower); err != nil {
        return utils.StatusError{http.StatusUnprocessableEntity, err}
	}
    rowsAffected, err := db.UpdateOne(nounSingularLower, id)
    if err != nil {
        return utils.StatusError{http.StatusInternalServerError, err}
    }
    w.Header().Set("Content-Type", "application/json")
    if *rowsAffected == 0 {
        w.WriteHeader(http.StatusNotFound)
    } else {
        w.WriteHeader(http.StatusOK)
    }
    return nil
}

func Delete<%= nounSingularUpper %>(db *database.Database, w http.ResponseWriter, r *http.Request) error {
    id := mux.Vars(r)["id"]
    rowsAffected, err := db.DeleteOne(id)
    if err != nil {
        return utils.StatusError{http.StatusInternalServerError, err}
    }
    w.Header().Set("Content-Type", "application/json")
    if *rowsAffected == 0 {
        w.WriteHeader(http.StatusNotFound)
    } else {
        w.WriteHeader(http.StatusNoContent)
    }
    return nil
}