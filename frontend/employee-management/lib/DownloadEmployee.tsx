export default async function DownloadEmployee (token: string) {
    const response = await fetch(`${process.env.NEXT_PUBLIC_AUTH_URL}/api/employees/download`, {
        method: "GET",
        mode: 'cors',
        headers: {
            authorization: `Bearer ${token}`,
        },
    });
    if (!response.ok) {
        throw new Error("Cannot get employees for some reason");
    }
    // Read the response as a blob
    const blob = await response.blob();

    return blob;
}