import React from "react";

export type UserState = {
	user: Record<string, any> | null;
	accessToken: string | null;
	refreshToken: string | null;
	appAccessToken: string | null;
	appRefreshToken: string | null;
};

export interface UserStateContextInterface {
	user: UserState;
	setUser: React.Dispatch<React.SetStateAction<UserState>>;
}

export const defaultUserState: UserState = {
	user: null,
	accessToken: null,
	refreshToken: null,
	appAccessToken: null,
	appRefreshToken: null,
};

export const UserContext = React.createContext<UserStateContextInterface>({
	user: defaultUserState,
	setUser: () => {},
});
