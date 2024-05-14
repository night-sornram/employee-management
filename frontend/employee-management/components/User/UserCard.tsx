import { UserJson } from "@/interface";

export default function UserCard ({userData} : {userData: UserJson}) {
    return (
        <div className="flex flex-row w-full bg-p-lightgray rounded shadow-md px-6 py-6">
            <div className="w-[20%] bg-p-darkgray rounded-xl">

            </div>
            <div className="w-[50%] flex flex-col pl-3">
                <h1 className="text-2xl">{userData.title}. {userData.name} {userData.lastname}</h1>
                <h1 className="text-xl">{userData.title}. {userData.name} {userData.lastname}</h1>
            </div>
            <div className="w-[30%] flex flex-col">
                <h1 className="text-p-gray">Gender: {userData.gender}</h1>
                <h1 className="text-p-gray">Role: {userData.role}</h1>
            </div>
        </div>
    );
}