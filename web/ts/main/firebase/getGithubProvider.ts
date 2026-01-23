import firebase from '@/firebase'

export default (): firebase.auth.GithubAuthProvider => {
    const provider = new firebase.auth.GithubAuthProvider()

    provider.addScope('read:user')
    provider.addScope('user:email')

    return provider
}
