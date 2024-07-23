import { defineStore } from "pinia";
import settings from '../settings.json'
const path = settings.url

export const useUsersStore = defineStore('UsersStore', {
    state: () => ({
        users: [],
        usersErrors: "",
        departments: [],
        employees: [],
        passDialog: false,
    }),
    actions:{
        async getUsers(token) {
            const res = await fetch(path+'users/'+token,{
                method: "GET",
            })
            const data = await res.json()
            if (res.status === 200) {
                this.users = data.users
            }
            console.log(this.users)
        },
        async createUser(fn, ln, sn, dp, login, pass, role, token) {
            const res = await fetch(path+'user', {
                method: "POST",
                body: JSON.stringify({
                    firstname: fn,
                    lastname: ln,
                    surname: sn,
                    department_id: dp,
                    login: login,
                    password: pass,
                    role: role,
                    token: token
                })
            })
            if(res.status === 200) {
                this.getUsers(token)    
            } else {
                const data = await res.json()
                this.usersErrors = data.message
            }
        },
        getEmployeeItems(dep) {
            console.log("all users:", this.users)
            console.log("department id:", dep)
            var y = []
            const employ = this.users.filter((el) => el.Role != "user" && el.departmentId === dep)
            for (let i = 0; i < employ.length; i++) {
                const x = {
                    title: employ[i].lastname + " " + employ[i].firstname[0] + "." + employ[i].surname[0] + ".",
                    value: employ[i].uuid
                }
                y.push(x)
           }
           console.log("employes:", y)
           return y
        },
        async deleteUser(userid, token) {
            console.log(userid)
            const res = await fetch(path+'userdel',{
                method: "POST",
                body: JSON.stringify({
                    uuid: userid,
                    token: token,
                })
            })
            if (res.status === 200) {
                this.getUsers(token)
            }
        },
        async getDepartments(token) {
            const res = await fetch(path+'departments',{
                method: "POST",
                body: JSON.stringify({
                    mode: "admin",
                    token: token,
                })
            })
            const data = await res.json()
            if (res.status === 200) {
                this.departments = data.departments
            }
        },
        async addDepartment(name, dowork, token) {
            const res = await fetch(path+'department',{
                method: "POST",
                body: JSON.stringify({
                    department_name: name,
                    department_dowork: dowork,
                    token: token
                })
            })
            if (res.status === 200) {
                this.getDepartments(token)
            }
        },
        async changeUserPassword(uuid, pass, token) {
            const res = await fetch(path+'passwd',{
                method: "POST",
                body: JSON.stringify({
                    uuid: uuid,
                    password: pass,
                    token: token
                })
            })
            if (res.status === 200) {
                this.passDialog = false
            }
        },
    }, 
})