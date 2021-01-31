# Title: `[Application/Service name]`

**Owner:** [Person to contact about this service]

**Team:** [Team that manages this service]

**Importance:** [Importance of the service; One of Critical, Important, Useful, Experimental]

**Submitted:** [Date this document was submitted for review]

**Status:** [Current state of the document; One of: Draft, Discuss, Agreed, Abandoned]

# 1. Overview

## 1.1 Objective

What does this application/service do?

## 1.2 Design Goals

Why is it needed? What are the high-level tasks this will perform?

## 1.3 Non-Goals

List anything that is specifically excluded. This can include items omitted from this version or permanently.

## 1.4 References

List of other documents related to this RFC

## 1.5 Assumptions

List of assumptions related to this RFC and for the project as a whole

## 1.6 Glossary

List of terms used throughout the document that might not be known to all readers

## 1.7

List of assumptions related to this RFC and for the project as a whole

# 2. Architecture / Solution

This section is best done with diagrams.

## 2.1 Overview

A general overiew of the solution from a technical perspective. Highlighting any significant technical decisions.

## 2.2 System Architecture

Where does this system fit in with the rest of the services of the company?

## 2.2 Interaction Diagrams

Map out how this system with interacts with resources and other services in the company.

## 2.3 Software Architecture

Show the main components/modules of the service and how they will interact with each other.

## 2.4 Deviations from the Paved Road

Highlight and justify any decisions that are vastly different from the usual choices for the company.

# 3. APIs

List the proposed public APIs for this service. This will include their URL, request and response formats.

# 4. Data

# 4.1 DB Schemas

List the proposed database tables. Each table should include the columns, indexes and significant query patterns.

# 4.2 Data Collection and Analysis

Outline any data that is collected and analyzed. You should specifically highlight any data that might be of interest to other teams.

# 5. Cross Functional Requirements

## 5.1 Security
Is there any significant security issues or concerns for this system? 

### 5.1.1 Fraud and Abuse
Are there ways for this system to be abused or used to perpetrate fraud?

### 5.1.2 Data privacy
Are there any data privacy issues that need to be addressed?
Is this service handling any PII? If so, how will this be dealt with?

### 5.1.3 Authentication & Authorization
How will users authenticate to this services?
Are there different levels of authorization?

## 5.2 High Availability
What is the plan to monitor and maintain the uptime of this service?
Is there an uptime SLA?

### 5.2.1 Monitoring & Alerting
How will you be monitoring this service?
How will you be receiving alerts?
What are the key metrics (business and technical) that you will monitor?

### 5.2.2 Logging
Where will your logs be available?
What logging strategy will you be using?
Will there be any significant information stored in your logs?

## 5.3 Scalability
How will this service scale? Horizontally? Vertically?
Which resources do you expect to need scaling first (CPU, memory, database)?

## 5.4 Resiliency
How will this service behave when one or more of the dependencies are down?
Can this service automatically recover when the dependency is restored?
Can the service automatically recover from transient issues?
Can this service be started without external resources like databases being available? 

## 5.5 Performance
Call out any expected performance issues and how you plan to deal with them.

## 5.6 Configuration
How will this service be configured? 
Can this configuration be changed while the service is running?

## 5.7 Test Plan
Highlight any special plans or requirements related to both automated and manual testing.

## 5.8 Rollout Plan
How will this service be rolled out?
Are there any special considerations required?
Will any of the features be controlled by feature flags or rolled out incrementally?

## 5.9 Infrastructure Plan
What infrastructure does this service require?
How do you plan to acquire this infrastructure?

## 5.10 Dependencies
Highlight any external dependencies (services, teams, etc.) that necessary for the success of this proposal.

## 5.11 Build vs Buy
Compare and contrast building this solution versus buying an existing solution.
Should include considerations related to cost, performance, knowledge acquisition, support, etc.

## 5.12 Internationalization and localization
Is localization needed? If so, how do you plan to address this requirement?
