import { UserJson } from "@/interface";

export default function UserCard ({userData} : {userData: UserJson}) {
    return (
        <div className="flex flex-row w-full bg-slate-300 rounded shadow-md px-6 py-6">
            <div className="w-[20%] bg-slate-700 rounded-xl">

            </div>
            <div className="w-[50%] flex flex-col pl-3">
                <h1 className="text-2xl">{userData.title_en}. {userData.first_name_en} {userData.last_name_en}</h1>
                <h1 className="text-xl">{userData.title_en}. {userData.first_name_en} {userData.last_name_en}</h1>
            </div>
            <div className="w-[30%] flex flex-col">
                <h1 className="text-slate-600">Gender: {userData.gender}</h1>
                <h1 className="text-slate-600">Role: {userData.role}</h1>
            </div>
        </div>
    );
}