export default async function DownloadAttendance (token: string) {
    const response = await fetch(`${process.env.NEXT_PUBLIC_ATTENDANCE_URL}/api/attendances/download`, {
        method: "GET",
        mode: 'cors',
        headers: {
            authorization: `Bearer ${token}`,
        },
    });
    if (!response.ok) {
        throw new Error("Cannot get attendances for some reason");
    }
    // Read the response as a blob
    const blob = await response.blob();

    return blob;
}