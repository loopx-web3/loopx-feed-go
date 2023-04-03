import * as components from "./loopxComponents"
export * from "./loopxComponents"

/**
 * @description "user register with wallet, email or social media"
 * @param params
 * @param headers
 */
export function login(params: components.UserLoginReqParams, headers: components.UserLoginReqHeaders) {
	return webapi.post<components.UserInfoReply>(`/api/v1/user/login`, params, headers)
}

/**
 * @description "user wallet cpatcha, return the sign info and nonce"
 * @param req
 */
export function captcha(req: components.CaptchaReq) {
	return webapi.get<components.CaptchaReply>(`/api/v1/user/captcha`, req)
}

/**
 * @description "user logout"
 */
export function logout() {
	return webapi.get<components.ReplyStatus>(`/api/v1/user/logout`)
}

/**
 * @description "Returns a specified wallet user's public information"
 * @param params
 */
export function account(params: components.AccountReqParams, walletAddress: string) {
	return webapi.get<components.UserInfoReply>(`/api/v1/user/${walletAddress}`, params)
}

/**
 * @description "Returns information about an authorized user."
 */
export function me() {
	return webapi.get<components.UserInfoReply>(`/api/v1/user/me`)
}

/**
 * @description "Updates the authenticating user's settings"
 * @param req
 */
export function settings(req: components.UserSettingReq) {
	return webapi.post<components.UserInfoReply>(`/api/v1/user/settings`, req)
}

/**
 * @description "user info avator upload"
 */
export function avatar() {
	return webapi.post<components.UserInfoReply>(`/api/v1/user/avatar`)
}

/**
 * @description "Returns a list of users the specified user ID is following"
 * @param params
 */
export function following(params: components.UserFollowingReqParams, user_id: string) {
	return webapi.post<components.UserFollowingReply>(`/api/v1/user/${user_id}/following`, params)
}

/**
 * @description "Returns a list of users who are followers of the specified user"
 * @param params
 */
export function follower(params: components.UserFollowerReqParams, id: string) {
	return webapi.post<components.UserFollowerReply>(`/api/v1/user/${user_id}/follower`, params)
}

/**
 * @description "Return the authenticating user's notication list"
 */
export function notification() {
	return webapi.get<components.NotificationReply>(`/api/v1/user/notification`)
}

/**
 * @description "Return information about a User’s liking."
 * @param params
 */
export function read(params: components.LikeReqParams, id: string, type: string) {
	return webapi.get<components.LikeReply>(`/api/v1/user/like/${id}/${type}`, params)
}

/**
 * @description "Allows a user ID to like a contract, token or feed"
 * @param params
 */
export function update(params: components.LikeReqParams, id: string, type: string) {
	return webapi.post<components.ReplyStatus>(`/api/v1/user/like/${id}/${type}`, params)
}

/**
 * @description "Allows a user ID to like a contract, token or feed"
 * @param params
 */
export function delete(params: components.LikeReqParams, id: string, type: string) {
	return webapi.delete<components.ReplyStatus>(`/api/v1/user/like/${id}/${type}`, params)
}

/**
 * @description "Returns list of feeds of a specified user ID in a period"
 * @param req
 */
export function list(req: components.FeedsListReq) {
	return webapi.get<components.FeedsListReply>(`/api/v1/feeds/list`, req)
}

/**
 * @description "Returns comment list the specified user ID"
 * @param params
 */
export function read(params: components.CommentListReqParams, topic_id: string, topic_type: string) {
	return webapi.get<components.CommentListReply>(`/api/v1/comment/${topic_id}/${topic_type}`, params)
}

/**
 * @description "Allows a user ID to comment a contract, token or feed"
 * @param params
 */
export function update(params: components.CommentListReqParams, topic_id: string, topic_type: string) {
	return webapi.post<components.ReplyStatus>(`/api/v1/comment/${topic_id}/${topic_type}`, params)
}

/**
 * @description "Allows a user ID to delete a comment"
 * @param params
 */
export function delete(params: components.CommentDeleteReqParams, comment_id: string) {
	return webapi.delete<components.ReplyStatus>(`/api/v1/comment/${comment_id}`, params)
}

/**
 * @description "Returns search info"
 * @param params
 */
export function search(params: components.SearchReqParams) {
	return webapi.get<components.SearchReply>(`/api/v1/search`, params)
}
