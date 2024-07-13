import { VerificationLevel, IDKitWidget } from "@worldcoin/idkit";
import type { ISuccessResult } from "@worldcoin/idkit";
import React from "react";

interface WorldIDVerificationProps {
    onSuccess: (result: ISuccessResult) => void;
    onError: (error: Error) => void;
}

export type VerifyReply = {
    code: string;
    detail: string;
  };

const WorldIDVerification: React.FC<WorldIDVerificationProps> = ({ onSuccess, onError }) => {
    if (!process.env.NEXT_PUBLIC_WLD_APP_ID || !process.env.NEXT_PUBLIC_WLD_ACTION) {
        throw new Error("World ID environment variables are not set!");
    }

    const handleProof = async (result: ISuccessResult) => {
        console.log("Proof received from IDKit:\n", JSON.stringify(result));
        const reqBody = {
            merkle_root: result.merkle_root,
            nullifier_hash: result.nullifier_hash,
            proof: result.proof,
            verification_level: result.verification_level,
            action: process.env.NEXT_PUBLIC_WLD_ACTION,
            signal: "",
            petition_id: 1,
        };
        console.log("Sending proof to backend for verification:\n", JSON.stringify(reqBody));
        try {
            const res: Response = await fetch("http://localhost:4242/petition/sign", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(reqBody),
            });
            console.log("Received response from backend:\n", res);
            const data: VerifyReply = await res.json();
            console.log("Received response from backend:\n", data);
            if (res.status === 200) {
                console.log("Successful response from backend:\n", data);
                onSuccess(result);

            }
            else if (res.status === 403) {
                console.log("Error response from backend:\n", data);
                onError(new Error("You have already signed this petition!"));
            } else {
                console.log("Error response from backend:\n", data);
                throw new Error(`Error code ${res.status} (${data.code}): ${data.detail}` ?? "Unknown error.");
            }
        } catch (error) {
            onError(error as Error);
        }
    };

    return (
        <IDKitWidget
            action={process.env.NEXT_PUBLIC_WLD_ACTION!}
            app_id={process.env.NEXT_PUBLIC_WLD_APP_ID as `app_${string}`}
            onSuccess={onSuccess}
            handleVerify={handleProof}
            verification_level={VerificationLevel.Orb}
        >
            {({ open }) => (
                <button className="border border-black rounded-md" onClick={open}>
                    <div className="mx-3 my-1">Sign with World ID</div>
                </button>
            )}
        </IDKitWidget>
    );
};

export default WorldIDVerification;
