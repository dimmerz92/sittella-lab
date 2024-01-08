import { post } from '../helpers';

export type GPTMessage = {
	role: string;
	content: string;
};

export async function sendGPTMessage(payload: GPTMessage[]) {
	try {
		const res: GPTMessage = await post('/api/chat', payload);
		return res;
	} catch (err) {
		console.log(err);
		return;
	}
}
