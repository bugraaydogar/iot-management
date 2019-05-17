// -*- Mode: Go; indent-tabs-mode: t -*-

/*
 * This file is part of the IoT Management Service
 * Copyright 2019 Canonical Ltd.
 *
 * This program is free software: you can redistribute it and/or modify it
 * under the terms of the GNU Affero General Public License version 3, as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful, but WITHOUT
 * ANY WARRANTY; without even the implied warranties of MERCHANTABILITY,
 * SATISFACTORY QUALITY, or FITNESS FOR A PARTICULAR PURPOSE.
 * See the GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package web

//// UserResponse is the response from a user creation/update
//type UserResponse struct {
//	Success      bool           `json:"success"`
//	ErrorCode    string         `json:"error_code"`
//	ErrorMessage string         `json:"message"`
//	User         datastore.User `json:"user"`
//}
//
//// UsersResponse is the response from a user list request
//type UsersResponse struct {
//	Success      bool             `json:"success"`
//	ErrorCode    string           `json:"error_code"`
//	ErrorMessage string           `json:"message"`
//	Users        []datastore.User `json:"users"`
//}

//
//func formatUserResponse(success bool, errorCode, message string, user datastore.User, w http.ResponseWriter) error {
//	response := UserResponse{Success: success, ErrorCode: errorCode, ErrorMessage: message, User: user}
//
//	// Encode the response as JSON
//	if err := json.NewEncoder(w).Encode(response); err != nil {
//		log.Println("Error forming the user response.")
//		return err
//	}
//	return nil
//}
//
//func formatUsersResponse(success bool, errorCode, message string, users []datastore.User, w http.ResponseWriter) error {
//	response := UsersResponse{Success: success, ErrorCode: errorCode, ErrorMessage: message, Users: users}
//
//	// Encode the response as JSON
//	if err := json.NewEncoder(w).Encode(response); err != nil {
//		log.Println("Error forming the users response.")
//		return err
//	}
//	return nil
//}
//
//// UsersHandler is the API method to list the users
//func (wb Service) UsersHandler(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", JSONHeader)
//
//	_, err := wb.checkIsSuperuserAndGetUserFromJWT(w, r)
//	if err != nil {
//		formatUsersResponse(false, "error-auth", "", nil, w)
//		return
//	}
//
//	users, err := datastore.Environ.DB.ListUsers()
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		formatUsersResponse(false, "error-fetch-users", err.Error(), nil, w)
//		return
//	}
//
//	// Return successful JSON response with the list of users
//	w.WriteHeader(http.StatusOK)
//	formatUsersResponse(true, "", "", users, w)
//}
//
//// UserCreateHandler handles user creation
//func (wb Service) UserCreateHandler(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", JSONHeader)
//
//	_, err := wb.checkIsSuperuserAndGetUserFromJWT(w, r)
//	if err != nil {
//		formatUsersResponse(false, "error-auth", "", nil, w)
//		return
//	}
//
//	user := datastore.User{}
//	err = json.NewDecoder(r.Body).Decode(&user)
//	switch {
//	// Check we have some data
//	case err == io.EOF:
//		w.WriteHeader(http.StatusBadRequest)
//		formatUserResponse(false, "error-user-data", "No user data supplied", user, w)
//		return
//		// Check for parsing errors
//	case err != nil:
//		w.WriteHeader(http.StatusBadRequest)
//		formatUserResponse(false, "error-data-json", err.Error(), user, w)
//		return
//	}
//
//	userID, err := datastore.Environ.DB.CreateUser(user)
//	if err != nil {
//		w.WriteHeader(http.StatusBadRequest)
//		formatUserResponse(false, "error-user-invalid", err.Error(), user, w)
//		return
//	}
//	user.ID = userID
//
//	formatUserResponse(true, "", "", user, w)
//}
//
//// UserGetHandler is the API method to retrieve user info
//func (wb Service) UserGetHandler(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", JSONHeader)
//
//	_, err := wb.checkIsSuperuserAndGetUserFromJWT(w, r)
//	if err != nil {
//		formatUserResponse(false, "error-auth", "", datastore.User{}, w)
//		return
//	}
//
//	vars := mux.Vars(r)
//	id, err := strconv.ParseInt(vars["id"], 10, 64)
//
//	if err != nil {
//		w.WriteHeader(http.StatusNotFound)
//		errorMessage := fmt.Sprintf("%v", vars)
//		formatUserResponse(false, "error-user-invalid", errorMessage, datastore.User{}, w)
//		return
//	}
//
//	user, err := datastore.Environ.DB.GetUser(id)
//	if err != nil {
//		w.WriteHeader(http.StatusNotFound)
//		errorMessage := fmt.Sprintf("User ID: %d.", id)
//		formatUserResponse(false, "error-get-user", errorMessage, datastore.User{ID: id}, w)
//		return
//	}
//
//	w.WriteHeader(http.StatusOK)
//	formatUserResponse(true, "", "", user, w)
//}
//
//// UserUpdateHandler handles user update
//func (wb Service) UserUpdateHandler(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", JSONHeader)
//
//	_, err := wb.checkIsSuperuserAndGetUserFromJWT(w, r)
//	if err != nil {
//		formatUsersResponse(false, "error-auth", "", nil, w)
//		return
//	}
//
//	user := datastore.User{}
//	err = json.NewDecoder(r.Body).Decode(&user)
//	switch {
//	// Check we have some data
//	case err == io.EOF:
//		w.WriteHeader(http.StatusBadRequest)
//		formatUserResponse(false, "error-user-data", "No user data supplied", user, w)
//		return
//		// Check for parsing errors
//	case err != nil:
//		w.WriteHeader(http.StatusBadRequest)
//		formatUserResponse(false, "error-data-json", err.Error(), user, w)
//		return
//	}
//
//	err = datastore.Environ.DB.UpdateUser(user)
//	if err != nil {
//		w.WriteHeader(http.StatusBadRequest)
//		formatUserResponse(false, "error-user-invalid", err.Error(), user, w)
//		return
//	}
//
//	formatUserResponse(true, "", "", user, w)
//}
