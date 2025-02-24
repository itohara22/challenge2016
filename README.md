# Distributor Permission Tool - Instructions

This tool allows you to manage distributors and their regional permissions via a command-line interface.

## Commands and Their Functionality

### 1. Creating a New Distributor
**Command:**
```sh
make <distributor_name>
```
**Example:**
```sh
make a
```
**Functionality:**
- Creates a new distributor named `a`.
- If the distributor already exists, an error message is displayed.

### 2. Creating a Distributor with a Parent
**Command:**
```sh
make <child_distributor> < parent_distributor
```
**Example:**
```sh
make b < a
```
**Functionality:**
- Creates a new distributor `b` and assigns `a` as its parent.
- If the parent does not exist, an error message is displayed.
- If the child distributor already exists, an error message is displayed.

### 3. Adding an Include Permission
**Command:**
```sh
for <distributor_name> include <place_code>,...
```
**Example:**
```sh
for a include IN
for a include IN,US
```
**Functionality:**
- Grants `a` access to `IN` (India) in first example.
- Use , for multiple code
- If the distributor does not exist, an error message is displayed.

### 4. Adding an Exclude Permission
**Command:**
```sh
for <distributor_name> exclude <place_code>,...
```
**Example:**
```sh
for a exclude AR-IN
for a exclude AR-IN,US
```
**Functionality:**
- Revokes access to `AR-IN` (Arunachal Pradesh, India) for `a` in first example. 
- Use , for multiple code
- If the distributor does not exist, an error message is displayed.

### 5. Listing Permissions
**Command:**
```sh
for <distributor_name> list
```
**Example:**
```sh
for a list
```
**Functionality:**
- Displays the include and exclude permissions for `a`.
- If the distributor does not exist, an error message is displayed.

### 6. Exiting the CLI
**Command:**
```sh
exit
```
**Functionality:**
- Terminates the CLI tool.

## Notes
- `place_code` can be:
  - `CountryCode` (e.g., `IN` for India)
  - `ProvinceCode-CountryCode` (e.g., `UP-IN` for Uttar Pradesh, India)
  - `CityCode-ProvinceCode-CountryCode` (e.g., `ZAIDR-UP-IN` for a specific city in Uttar Pradesh, India)

## Example Workflow
```sh
make a
for a include IN
for a exclude AR-IN
make b < a
for a list
for b list
exit
```
### Expected Output:
```
Created distributor: a
a now includes IN
a now excludes AR-IN
Created distributor: b with parent: a
Permissions for a:
Includes: map[IN:true]
Excludes: map[AR-IN:true]
Permissions for b:
Includes: map[]
Excludes: map[]
Exiting CLI...
```
