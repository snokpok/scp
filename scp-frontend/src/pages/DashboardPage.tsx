import React from "react";
import { useNavigate } from "react-router-dom";
import Cookies from "universal-cookie/es6";
import { UserContext } from "../common/contexts/user.context";
import { SERVER_URL } from "../common/serverqueries";
import LogoutButton from "../components/Auth/LogoutButton";
import SCPWidget from "../components/User/SCPWidget";
import UserWidget from "../components/User/UserWidget";

const cookies = new Cookies();

function DashboardPage() {
	const { user, setUser } = React.useContext(UserContext);
	const navigate = useNavigate();
	const scpURLApi = `${SERVER_URL}/scp`;

	React.useEffect(() => {
		const appAccessToken = cookies.get("accessToken") ?? user.appAccessToken;
		if (!appAccessToken) {
			navigate("/", {
				replace: true,
			});
		} else {
			// try to fetch all info with this
			setUser((prev) => ({
				...prev,
				appAccessToken,
			}));
		}
	}, [navigate, user.appAccessToken, setUser]);

	return (
		<div className="flex flex-col bg-green-600 w-screen min-h-screen items-center justify-center space-y-2">
			<div className="flex flex-col items-center space-y-2">
				<LogoutButton />
				<UserWidget />
			</div>
			<div className="flex flex-col items-center justify-center rounded-lg w-96 p-5 bg-white">
				<h1>
					<div className="font-bold text-lg">Welcome!</div>
				</h1>
				<div>You can access your currently playing song by this URL:</div>
				<input
					value={scpURLApi}
					readOnly
					className="p-2 bg-gray-300 rounded-sm"
				/>
				<br />
				<label htmlFor="access-token-spotify">Access token spotify</label>
				<input
					value={user.accessToken ?? ""}
					readOnly
					className="p-2 bg-gray-300 rounded-sm"
					id="access-token-spotify"
				/>
				<label htmlFor="refresh-token-spotify">Refresh token spotify</label>
				<input
					value={user.refreshToken ?? ""}
					readOnly
					className="p-2 bg-gray-300 rounded-sm"
					id="refresh-token-spotify"
				/>
				<br />
				<label htmlFor="access-token-app">
					Access token for this app (use this)
				</label>
				<input
					value={user.appAccessToken ?? ""}
					readOnly
					className="p-2 bg-gray-300 rounded-sm"
					id="access-token-app"
				/>
			</div>
			<div className="bg-black p-2">
				<h2 className="text-white">Request with cURL:</h2>
			</div>
			<code className="p-2 bg-gray-700 text-white max-w-2xl overflow-x-scroll">
				curl --location --request GET '{scpURLApi}' {"\n"} --header
				'Authorization: Bearer {user.appAccessToken}'
			</code>
			<div className="p-2 bg-black space-y-2 flex flex-col items-center rounded-md">
				<div>
					<h1 className="font-bold text-white">
						Sample custom widget from fetched data:
					</h1>
				</div>
				<SCPWidget />
			</div>
		</div>
	);
}

export default DashboardPage;
