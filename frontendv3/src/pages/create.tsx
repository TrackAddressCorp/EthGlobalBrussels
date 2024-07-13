// pages/index.tsx
import { useState, ChangeEvent } from 'react';
import {
	useToast,
	Button,
	Input,
	Textarea,
	Box,
	FormControl,
	FormLabel,
	ChakraProvider,
} from "@chakra-ui/react";
import { useRouter } from 'next/navigation';

export default function Home() {
	const router = useRouter();

	const [title, setTitle] = useState<string>('');
	const [description, setDescription] = useState<string>('');
	const [petitionId, setPetitionId] = useState<number | null>(null);
	const [file, setFile] = useState<File | null>(null);
	const toast = useToast();
	const [uploading, setLoading] = useState(false);

	const handleCreatePetition = async (): Promise<void> => {
		const response = await fetch('http://localhost:4242/petition/create', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
			},
			body: JSON.stringify({ title, description }),
		});

		const data = await response.json();
		if (data.status_code === 200) {
			setPetitionId(data.id);
			toast({
				title: "Petition created.",
				description: data.status_msg,
				status: "success",
				duration: 5000,
				isClosable: true,
			});
		} else {
			toast({
				title: "Error.",
				description: "Failed to create petition.",
				status: "error",
				duration: 5000,
				isClosable: true,
			});
		}
	};

	const handleFileChange = (e: ChangeEvent<HTMLInputElement>) => {
		const file = e.target.files?.[0] || null;
		setFile(file);
	};

	const handleUploadPdf = async () => {
		if (!file || !petitionId) return;

		const formData = new FormData();
		formData.append('pdf', file);
		formData.append('petition_id', petitionId.toString());

		setLoading(true);

		const response = await fetch('http://localhost:4242/petition/upload', {
			method: 'POST',
			body: formData,
		});

		const data = await response.json();
		if (data.status_code === 200) {
			toast({
				title: "PDF uploaded.",
				description: "Your PDF has been uploaded successfully.",
				status: "success",
				duration: 5000,
				isClosable: true,
			});
		} else {
			toast({
				title: "Error.",
				description: "Failed to upload PDF.",
				status: "error",
				duration: 5000,
				isClosable: true,
			});
		}
		setLoading(false);
	};

	const handleFinishPetition = async () => {
		if (!petitionId) return;

		const response = await fetch(`http://localhost:4242/petition/finish/${petitionId}`, {
			method: 'POST',
		});

		const data = await response.json();
		if (data.status_code === 200) {
			toast({
				title: "Petition finished.",
				description: data.status_msg,
				status: "success",
				duration: 5000,
				isClosable: true,
			});
			router.push('/petition/' + petitionId);
		} else {
			toast({
				title: "Error.",
				description: "Failed to finish petition.",
				status: "error",
				duration: 5000,
				isClosable: true,
			});
		}
	};

	return (
		<ChakraProvider>
			<Box maxW="md" mx="auto" mt="10">
				<FormControl id="title" mb="4">
					<FormLabel>Title</FormLabel>
					<Input
						type="text"
						value={title}
						onChange={(e) => setTitle(e.target.value)}
					/>
				</FormControl>
				<FormControl id="description" mb="4">
					<FormLabel>Description</FormLabel>
					<Textarea
						value={description}
						onChange={(e) => setDescription(e.target.value)}
					/>
				</FormControl>

				{!petitionId && (
					<Button
						colorScheme="blue"
						onClick={handleCreatePetition}
						mb="4"
					>
						Create Petition
					</Button>
				)}

				{petitionId && (
					<div>
						<FormControl id="pdf" mb="4" >
							<FormLabel>Upload PDF</FormLabel>
							<Input
								type="file"
								accept="application/pdf"
								onChange={handleFileChange}
								style={{
									height: '50px', // Adjust height as needed
									alignContent: 'center',
								}}
							/>
						</FormControl>
						<div style={{ display: 'flex', flexDirection: 'column', gap: '10px' }}>
							<Button
								isLoading={uploading}
								colorScheme="green"
								onClick={handleUploadPdf}
								mb="4"
							>
								Upload PDF
							</Button>
							<Button
								colorScheme="red"
								mb="4"
								onClick={handleFinishPetition}
								style={
									{
										alignItems: 'right',
									}

								}
							>
								Finish Petition
							</Button>
						</div>
					</div>
				)}
			</Box>
		</ChakraProvider>
	);
}
