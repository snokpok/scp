import React from "react";
import Cookies from "universal-cookie/es6";
import {
	defaultUserState,
	UserContext,
} from "../../common/contexts/user.context";

const cookies = new Cookies();

function LogoutButton() {
	const { setUser } = React.useContext(UserContext);
	const handleLogout = () => {
		cookies.remove("accessToken");
		setUser(() => defaultUserState);
	};

	return (
		<button
			className="bg-black p-2 rounded-md hover:bg-red-500"
			onClick={(e) => {
				e.preventDefault();
				handleLogout();
			}}
		>
			<p className="text-white">Logout</p>
		</button>
	);
}

export default LogoutButton;
