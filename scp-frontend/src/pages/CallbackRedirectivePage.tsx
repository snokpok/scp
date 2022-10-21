import React from "react";
import { useNavigate } from "react-router-dom";
import { UserContext } from "../common/contexts/user.context";
import { getMeFromSpotify, requestToken } from "../common/queryfunctions";
import { addUser } from "../common/serverqueries";

function CallbackRedirectivePage() {
	const { user, setUser } = React.useContext(UserContext);
	const navigate = useNavigate();
	const [startRedirect, setStartRedirect] = React.useState(false);
	let query = React.useMemo(
		() => new URLSearchParams(window.location.search),
		[]
	);

	const handleQueryParseAuth = React.useCallback(() => {
		if (!query.get("code")) {
			alert("please accept the permissions");
			navigate("/", {
				replace: true,
			});
		} else {
			requestToken(query.get("code") as string).then((res) => {
				const accessToken = res.data["access_token"];
				const refreshToken = res.data["refresh_token"];
				getMeFromSpotify(accessToken).then(({ data: myData }) => {
					if (accessToken && refreshToken) {
						const username = myData["display_name"];
						const email = myData["email"];
						const spotifyId = myData["id"];
						addUser({
							username,
							email,
							access_token: accessToken,
							refresh_token: refreshToken,
							spotify_id: spotifyId,
						})
							.then(({ data }) => {
								const appAccessToken = data["data"]["token"];
								setUser((prev) => ({
									...prev,
									appAccessToken,
								}));
								document.cookie = `accessToken=${appAccessToken}; path=/`;
							})
							.catch((reason) => {
								alert(reason);
								navigate("/", {
									replace: true,
								}); // redirect to login page if any error
							});
						setStartRedirect(true);
					} else {
						alert(
							"something went wrong with the access token retrieval process"
						);
					}
				});
			});
		}
	}, [query, navigate, setUser]);

	React.useEffect(() => {
		if (user.appAccessToken && startRedirect)
			navigate("/dashboard", {
				replace: true,
			});
	}, [startRedirect, navigate, user.appAccessToken]);

	React.useEffect(() => {
		handleQueryParseAuth();
	}, [handleQueryParseAuth]);

	return (
		<div>
			<div>Wait a bit...</div>
		</div>
	);
}

export default CallbackRedirectivePage;
