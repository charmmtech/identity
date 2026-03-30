# Identity Service Architecture

This service is a stateless bridge to an upstream identity provider.

Current provider:
- Keycloak admin API via `gocloak`

Explicit non-goals:
- No application-owned user database
- No persistence layer for identities, roles, groups, sessions, or MFA state

The protobuf surface is intentionally provider-oriented instead of Keycloak-shaped:
- `User`, `Role`, `Group`, `Session`, and `MfaStatus` are generic transport models
- `ProviderRef` identifies the backing provider type and realm
- RPCs expose CRUD and assignment flows that should remain stable if the implementation moves from Keycloak to Ory later

Current RPC coverage:
- User CRUD
- Enable/disable user
- Reset/set user password
- Assign/remove realm roles to users
- Add/remove users from groups
- List/revoke user sessions
- Read/update MFA state
- Role CRUD
- Group CRUD
- Assign/remove realm roles to groups

Keycloak-specific notes:
- MFA support currently maps to TOTP state, required actions, and credential disablement
- Group membership responses return group IDs on users to avoid assuming provider-specific path semantics
- Role operations currently target realm roles, not client roles
