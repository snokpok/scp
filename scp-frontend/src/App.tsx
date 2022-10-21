import React from "react";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import { UserContext, UserState } from "./common/contexts/user.context";
import CallbackRedirectivePage from "./pages/CallbackRedirectivePage";
import DashboardPage from "./pages/DashboardPage";
import LoginPage from "./pages/LoginPage";

function App() {
	const [user, setUser] = React.useState<UserState>({
		user: null,
		accessToken: null,
		refreshToken: null,
		appAccessToken: null,
		appRefreshToken: null,
	});

	return (
		<div className="h-screen w-screen">
			<UserContext.Provider value={{ user, setUser }}>
				<BrowserRouter>
					<Routes>
						<Route path="/" element={<LoginPage />} />
						<Route path="/callback" element={<CallbackRedirectivePage />} />
						<Route path="/dashboard" element={<DashboardPage />} />
					</Routes>
				</BrowserRouter>
			</UserContext.Provider>
		</div>
	);
}

export default App;
