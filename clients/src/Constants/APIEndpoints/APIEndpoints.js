export default {
    base: "https://studybuddy-api.kaylalee.me",
    testbase: "https://localhost:4000",
    handlers: {
        users: "/students",
        myuser: "/students/me",
        myuserAvatar: "/students/me/avatar",
        sessions: "/sessions",
        sessionsMine: "/sessions/mine",
        resetPasscode: "/resetcodes",
        passwords: "/passwords/"
    }
}