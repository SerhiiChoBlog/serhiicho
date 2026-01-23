import firebase from '@/firebase'

export default (): firebase.auth.GoogleAuthProvider => {
    const provider = new firebase.auth.GoogleAuthProvider()
    provider.addScope('email')
    provider.addScope('profile')

    return provider
}
