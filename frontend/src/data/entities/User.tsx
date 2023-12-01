export interface IUser {
    id: number
    type: string
    name: string
    username: string
    email: string
    photo : string
    phone: string
    roleId: number
    preferredUsername: string
    summary: string
}

class User implements IUser {
    id: number
    type: string
    name: string
    username: string
    email: string
    photo : string
    phone: string
    roleId: number
    preferredUsername: string
    summary: string

    constructor(props: IUser) {
        this.id = props.id
        this.type = props.type
        this.name = props.name
        this.username = props.username
        this.email = props.email
        this.photo = props.photo
        this.phone = props.phone
        this.roleId= props.roleId
        this.preferredUsername = props.preferredUsername
        this.summary = props.summary
    }
}

const createUser = (props: IUser): User => new User(props)

const initalUser = (): User => {
    return createUser({
        id: 0,
        type: '',
        name: '',
        username: '',
        email: '',
        photo : '',
        phone: '',
        roleId: 0,
        preferredUsername: '',
        summary: '',
    })
}
export { createUser, initalUser, User }
