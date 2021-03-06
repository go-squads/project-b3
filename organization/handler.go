package organization

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/go-squads/reuni-server/appcontext"

	"github.com/go-squads/reuni-server/helper"
	"github.com/gorilla/mux"

	"github.com/go-squads/reuni-server/response"
)

var proc processor

func getProcessor() processor {
	if proc == nil {
		proc = &mainProcessor{repo: initRepository(appcontext.GetHelper())}
	}
	return proc
}

func getFromContext(r *http.Request, key string) string {
	data := r.Context().Value(key)
	if data == nil {
		return ""
	}
	return fmt.Sprintf("%v", data)
}

func CreateOrganizationHandler(w http.ResponseWriter, r *http.Request) {
	var data Organization
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response.ResponseError("CreateOrganization", getFromContext(r, "username"), w, helper.NewHttpError(http.StatusBadRequest, err.Error()))
		return
	}
	if data.Name == "" {
		response.ResponseError("CreateOrganization", getFromContext(r, "username"), w, helper.NewHttpError(http.StatusBadRequest, "Name cannot be empty"))
		return
	}
	reg, _ := regexp.Compile(`^[^.|\s]+$`)
	if !reg.MatchString(data.Name) {
		response.ResponseError("CreateOrganization", getFromContext(r, "username"), w, helper.NewHttpError(http.StatusBadRequest, "Organization name should not contain '.' or any whitespaces"))
		return
	}
	uid, err := strconv.ParseInt(getFromContext(r, "userId"), 10, 64)
	if err != nil {
		response.ResponseError("CreateOrganization", getFromContext(r, "username"), w, helper.NewHttpError(http.StatusInternalServerError, err.Error()))
		return
	}
	err = getProcessor().createNewOrganizationProcessor(data.Name, uid)
	if err != nil {
		response.ResponseError("CreateOrganization", getFromContext(r, "username"), w, err)
		return
	}
	defer r.Body.Close()
	response.ResponseHelper(w, http.StatusCreated, response.ContentText, "201 Created")
}

func AddUserHandler(w http.ResponseWriter, r *http.Request) {
	var member Member
	organizationName := mux.Vars(r)["organization_name"]
	organizationId, err := getProcessor().translateNameToIdProcessor(organizationName)

	if err != nil {
		response.ResponseError("AddUser", getFromContext(r, "username"), w, helper.NewHttpError(http.StatusInternalServerError, err.Error()))
		return
	}
	log.Printf("AddUser: Get Request to add user to org: %d", organizationId)
	err = json.NewDecoder(r.Body).Decode(&member)
	defer r.Body.Close()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Add user: error parsing body")
		return
	}
	member.OrgId = int64(organizationId)
	err = getProcessor().addUserProcessor(&member)
	if err != nil {
		response.ResponseError("AddUser", getFromContext(r, "username"), w, err)
		return
	}
	response.ResponseHelper(w, http.StatusCreated, response.ContentText, "201 Created")
}

func DeleteUserFromGroupHandler(w http.ResponseWriter, r *http.Request) {
	var member Member
	organizationName := mux.Vars(r)["organization_name"]
	organizationId, err := getProcessor().translateNameToIdProcessor(organizationName)
	orgId := int64(organizationId)
	if err != nil {
		response.ResponseError("DeleteUser", getFromContext(r, "username"), w, helper.NewHttpError(http.StatusInternalServerError, err.Error()))
		return
	}
	log.Printf("DeleteUser: Get Request to delete user to org: %d", organizationId)
	err = json.NewDecoder(r.Body).Decode(&member)
	defer r.Body.Close()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("DeleteUser: error parsing body")
		return
	}
	err = getProcessor().deleteUserFromGroupProcessor(orgId, member.UserId)
	if err != nil {
		response.ResponseError("DeleteUser", getFromContext(r, "username"), w, err)
		return
	}
	response.ResponseHelper(w, http.StatusOK, response.ContentText, "200 OK")
}

func UpdateRoleOfUserHandler(w http.ResponseWriter, r *http.Request) {
	var member Member
	organizationName := mux.Vars(r)["organization_name"]
	organizationId, err := getProcessor().translateNameToIdProcessor(organizationName)
	orgId := int64(organizationId)
	if err != nil {
		response.ResponseError("AddUser", getFromContext(r, "username"), w, helper.NewHttpError(http.StatusInternalServerError, err.Error()))
		return
	}
	log.Printf("UpdateRoleOfUser: Get Request to update user %d role to: %s", orgId, member.Role)
	err = json.NewDecoder(r.Body).Decode(&member)
	defer r.Body.Close()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("UpdateRoleOfUser: error parsing body")
		return
	}
	member.OrgId = orgId
	err = getProcessor().updateRoleOfUserProcessor(&member)
	if err != nil {
		response.ResponseError("UpdateRoleOfUser", getFromContext(r, "username"), w, err)
		return
	}
	response.ResponseHelper(w, http.StatusOK, response.ContentText, "200 OK")

}

func GetAllMemberOfOrganizationHandler(w http.ResponseWriter, r *http.Request) {
	organizationName := mux.Vars(r)["organization_name"]
	organizationId, err := getProcessor().translateNameToIdProcessor(organizationName)
	orgID := int64(organizationId)
	if err != nil {
		response.ResponseError("GetAllMemberOfOrganization", getFromContext(r, "username"), w, helper.NewHttpError(http.StatusInternalServerError, err.Error()))
		return
	}
	fmt.Println("GETALLMEMBER of orgID: " + fmt.Sprint(orgID))
	members, err := getProcessor().getAllMemberOfOrganizationProcessor(orgID)
	if err != nil {
		response.ResponseError("GetAllMemberOfOrganization", getFromContext(r, "username"), w, err)
		return
	}
	response.ResponseHelper(w, http.StatusOK, response.ContentJson, string(members))
}

func GetAllHandler(w http.ResponseWriter, r *http.Request) {
	uid, err := strconv.ParseInt(getFromContext(r, "userId"), 10, 64)
	if err != nil {
		response.ResponseError("GetAllHandler", getFromContext(r, "username"), w, helper.NewHttpError(http.StatusBadRequest, "Cannot get User ID"))
		return
	}
	orgs, err := getProcessor().getAllOrganizationProcessor(int(uid))

	if err != nil {
		response.ResponseError("GetAllHandler", getFromContext(r, "username"), w, err)
		return
	}
	response.ResponseHelper(w, http.StatusOK, response.ContentJson, orgs)
}
