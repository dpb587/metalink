// For metalink repositories stored in a git repository.
//
// **Supported schemes**
//
//  * `git`
//  * `git+file`
//  * `git+http`
//  * `git+https`
//  * `git+ssh`
//
// **URI components**
//
// *Path* - split on a double forward slash. The first part will be used as the path to `git clone`, and all remaining
// parts will be used as a subpath within the repository.
//
// *Fragment* - a branch to clone
//
// **Example URIs**
//
// Use a non-default branch of `metalink-repository`...
//
//     git+https://github.com/acmecorp/prod-assets.git#metalink-repository
//
// Use a subpath within a repository...
//
//     git+ssh://git.example.com/prod.git//assets/metalink-repository
//
// **Notes**
//
//  * for SSH-based connections, the user's default SSH configuration/key/agent is used
package git
