import { FaSpotify } from "react-icons/fa";
import { originalCallbackURL } from "../../common/data";
import { CLIENT_ID } from "../../common/env";

function LoginForm() {
	const stateNumber = "34fFs29kd09";
	const scopes =
		"user-read-currently-playing user-read-private user-read-email";
	const authorizeSpotifyURL = `https://accounts.spotify.com/authorize?client_id=${CLIENT_ID}&response_type=code&redirect_uri=${encodeURIComponent(
		originalCallbackURL
	)}&scope=${scopes}&state=${stateNumber}`;

	return (
		<a href={authorizeSpotifyURL}>
			<div className="flex flex-col justify-center items-center bg-white text-white rounded-lg">
				<div className="flex bg-black p-4 rounded-md space-x-1 cursor-pointer">
					<div className="flex justify-center items-center bg-black">
						<FaSpotify className="text-green-400" />
					</div>
					<div className="font-bold">Login via Spotify</div>
				</div>
			</div>
		</a>
	);
}

export default LoginForm;
