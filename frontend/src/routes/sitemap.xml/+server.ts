export async function GET() {
	return new Response(
		`
    <?xml version="1.0" encoding="UTF-8"?>
    <urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
      <url>
        <loc>https://www.sittellalab.com.au/</loc>
        <lastmod>2024-01-12</lastmod>
      </url>
    </urlset>
    `.trim(),
		{
			headers: {
				'Content-Type': 'application/xml'
			}
		}
	);
}
