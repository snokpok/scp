import axios, { AxiosResponse } from "axios";
import qs from "querystring";
import { originalCallbackURL } from "./data";
import { CLIENT_ID, CLIENT_SECRET } from "./env";

export const getClientHeaderField = (clientId: string, clientSecret: string) =>
	Buffer.from(`${clientId}:${clientSecret}`).toString("base64");

const headerField = getClientHeaderField(CLIENT_ID, CLIENT_SECRET);

export const requestToken = async (code: string) => {
	const res = axios({
		method: "POST",
		url: "https://accounts.spotify.com/api/token",
		headers: {
			Authorization: `Basic ${headerField}`,
			Accept: "application/json",
			"content-type": "application/x-www-form-urlencoded",
		},
		data: qs.stringify({
			grant_type: "authorization_code",
			code,
			redirect_uri: originalCallbackURL,
		}),
	});
	return res as unknown as AxiosResponse<any>;
};

export async function refreshToken(
	refreshToken: string
): Promise<AxiosResponse<any>> {
	const res = axios({
		url: "https://accounts.spotify.com/api/token",
		headers: {
			Authorization: `Basic ${headerField}`,
			Accept: "application/json",
			"content-type": "application/x-www-form-urlencoded",
		},
		data: qs.stringify({
			grant_type: "refresh_token",
			refresh_token: refreshToken,
		}),
	});
	return res;
}

export async function getMeFromSpotify(
	accessToken: string
): Promise<AxiosResponse<any>> {
	const res = axios({
		url: "https://api.spotify.com/v1/me",
		headers: {
			Authorization: `Bearer ${accessToken}`,
			Accept: "application/json",
		},
	});
	return res;
}
