import React from "react";
import { UserContext } from "../../common/contexts/user.context";
import { getMeFromServer } from "../../common/serverqueries";

function UserWidget() {
	const { user, setUser } = React.useContext(UserContext);

	React.useEffect(() => {
		// fetch the user info with access token from data
		// if no access token then fetch the user from our db then update
		if (!user.accessToken && user.appAccessToken) {
			getMeFromServer(user.appAccessToken).then(({ data }) => {
				setUser((prev) => ({
					...prev,
					accessToken: data.access_token,
					refreshToken: data.refresh_token,
					user: {
						display_name: data.username,
						email: data.email,
						spotify_id: data.spotify_id,
					},
				}));
			});
		}
	}, [user, setUser]);

	if (!user || !user.user) {
		return <div>Loading...</div>;
	}

	return (
		<div className="flex rounded-md bg-white p-2 min-h-20 space-x-2">
			<div className="flex flex-col text-center">
				<p>
					<div className="font-bold">
						<h4>{user.user["display_name"]}</h4>
					</div>
				</p>
				<p>
					<div className="text-gray-600 italic">{user.user["email"]}</div>
				</p>
				<p>
					<div className="text-gray-400 italic text-xs">
						{user.user["spotify_id"]}
					</div>
				</p>
			</div>
		</div>
	);
}

export default UserWidget;
