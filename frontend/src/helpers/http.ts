async function http<T>(req: Request): Promise<T> {
	console.log(req);
	const response = await fetch(req);
	if (!response.ok) throw new Error(await response.text());

	const headers = response.headers;
	const data = headers.get('Content-Type')?.includes('json') ? await response.json() : {};

	return data;
}

async function request<Req>(path: string, method: string, body?: Req): Promise<Request> {
	return new Request(path, {
		method,
		headers: {
			'Content-Type': 'application/json'
		},
		body: JSON.stringify(body)
	});
}

export async function get<Res>(path: string): Promise<Res> {
	const req = await request(path, 'GET');
	return await http<Res>(req);
}

export async function post<Req, Res>(path: string, body: Req): Promise<Res> {
	const req = await request<Req>(path, 'POST', body);
	return await http<Res>(req);
}
