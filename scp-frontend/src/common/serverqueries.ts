import axios from "axios";
import { SERVER_URL } from "./env";

export interface AddUserArgs {
	username: string;
	email: string;
	spotify_id: string;
	access_token: string;
	refresh_token: string;
}

export async function addUser(user: AddUserArgs) {
	return axios({
		method: "POST",
		url: `${SERVER_URL}/user`,
		data: user,
	});
}

export async function getSCPFromServer(accessToken: string) {
	return axios({
		method: "GET",
		headers: {
			Authorization: `Bearer ${accessToken}`,
		},
		url: `${SERVER_URL}/scp`,
	});
}

export async function getMeFromServer(accessToken: string) {
	return axios({
		method: "GET",
		headers: {
			Authorization: `Bearer ${accessToken}`,
		},
		url: `${SERVER_URL}/me`,
	});
}
