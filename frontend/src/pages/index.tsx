import React from "react";
import WorldIDVerification from "./components/WorldIDVerificationProps"; // Adjust the path as needed
import type { ISuccessResult } from "@worldcoin/idkit";

const Home: React.FC = () => {
	const handleSuccess = (result: ISuccessResult) => {
		// This is where you should perform frontend actions once a user has been verified, such as redirecting to a new page
		window.alert("Successfully verified with World ID! Your nullifier hash is: " + result.nullifier_hash);
	};

	const handleError = (error: Error) => {
		console.error("Verification error:", error);
		window.alert("Verification failed: " + error.message);
	};

	return (
		<div>
			<div className="flex flex-col items-center justify-center align-middle h-screen">
				<p className="text-2xl mb-5">World ID Cloud Template</p>
				<WorldIDVerification onSuccess={handleSuccess} onError={handleError} />
			</div>
		</div>
	);
};

export default Home;
