export type GPTMessage = {
	role: string;
	content: string;
};

export async function sendGPTMessage(data: GPTMessage[]): Promise<GPTMessage | null> {
	try {
		const res = await fetch(`${import.meta.env.VITE_API_URL}/api/chat`, {
			method: 'POST',
			body: JSON.stringify(data)
		});

		if (!res.ok) throw new Error(`${res.status}: ${res.statusText}`);

		const reply: GPTMessage = await res.json();
		return reply;
	} catch (err) {
		console.log(err);
		return null;
	}
}
