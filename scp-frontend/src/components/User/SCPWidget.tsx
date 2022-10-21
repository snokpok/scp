import React from "react";
import { UserContext } from "../../common/contexts/user.context";
import { getSCPFromServer } from "../../common/serverqueries";

function SCPWidget() {
	const { user } = React.useContext(UserContext);
	const [scp, setScp] = React.useState<Record<string, any>>({});

	const fetchSCP = React.useCallback(() => {
		if (user.appAccessToken) {
			getSCPFromServer(user.appAccessToken).then(({ data }) => {
				setScp(data);
			});
		}
	}, [user.appAccessToken]);

	React.useEffect(() => {
		fetchSCP();
	}, [fetchSCP]);

	const RefetchButton = () => (
		<button
			onClick={() => {
				fetchSCP();
			}}
			className="bg-black p-2 text-white rounded-lg"
		>
			Refetch
		</button>
	);

	if (!scp) {
		return (
			<div className="bg-white rounded-lg p-2 flex items-center space-x-4 w-max">
				<div>No track!</div>
				<RefetchButton />
			</div>
		);
	}

	if (!scp["item"]) return <div>Loading...</div>;

	return (
		<div className="bg-white rounded-lg p-2 flex items-center space-x-4 w-max">
			<div className="rounded-full flex items-center space-x-2">
				<img
					src={scp["item"].album.images[2].url}
					className="rounded-full border-2 animate-spin"
					alt="Album cover"
				/>
				<div>
					<div className="font-bold">{scp["item"]?.name}</div>
					<div>
						{scp["item"]?.artists.map((item: any) => item.name).join(", ")}
					</div>
				</div>
			</div>
			<RefetchButton />
		</div>
	);
}

export default SCPWidget;
