import React from "react";
import { useNavigate } from "react-router-dom";
import Cookies from "universal-cookie/es6";
import { UserContext } from "../common/contexts/user.context";
import LoginForm from "../components/Auth/LoginForm";

const cookies = new Cookies();

function LoginPage() {
	const { setUser } = React.useContext(UserContext);
	const navigate = useNavigate();
	React.useEffect(() => {
		const appAccessToken = cookies.get("accessToken");
		if (appAccessToken) {
			setUser((prev) => ({
				...prev,
				appAccessToken,
			}));
			navigate("/dashboard", {
				replace: true,
			});
		}
	}, [navigate, setUser]);

	return (
		<div className="flex flex-col justify-center items-center bg-green-500 w-screen min-h-screen text-black space-y-2">
			<div className="flex flex-col text-center p-2 bg-white rounded-md">
				<h1 className="font-mono text-3xl font-bold">
					Spotify Currently Playing
				</h1>
				<p>
					This app exposes an API endpoint for you to get data fetched from the{" "}
					<code>/currently-playing</code> endpoint from the official Spotify API
				</p>
			</div>
			<LoginForm />
		</div>
	);
}

export default LoginPage;
