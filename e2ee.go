/* Copyright 2019 Sumukha PK
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

 package gomatrixserverlib

 import (
	 "context"
	 "net/url"
 )
 
 // KeyForwarding enables the client HS to talk to another clients HS.
 func (ac *FederationClient) KeyForwarding(
	 ctx context.Context, s ServerName, roomID, userID, endPoint string,
 ) (res RespMakeJoin, err error) {
	 path := federationPathPrefix + "/" + endPoint + "/" +
		 url.PathEscape(roomID) + "/" +
		 url.PathEscape(userID)
	 req := NewFederationRequest("GET", s, path)
	 err = ac.doRequest(ctx, req, &res)
	 return
 }
 