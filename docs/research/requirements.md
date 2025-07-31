# Requirements

---

1. Profile File -> `Yaml`
2. Profile Proxy.IPAddr -> HTTP
3. Generate Random Name for Profile
4. Get Random Browser, Device, Platform for Profile
5. Get Random Screen Size for Profile
6. Generate Firefox Profile Folder Name
7. Firefox Profile Management: Generate, Save, Update, Delete, Start, Search, Bulk Operations, etc.
8. HTTP Header `User-Agent` Parsing
9. Run Command `Firefox.exe`

---

- Generate Basic Files -- <https://ffprofile.com/>
  - pref.js -> firefox.cfg
  - policies.json

- Start New Profile -> Customize

- Config With Firefox autoconfig
  - FIREFOX-INSTALL-FOLDER
    - defaults/pref/autoconfig.js
    - firefox.cfg
    - distribution/policies.json
  - PROFILE-FOLDER
    - pref.js
    - extensions/***.xpi
      - chameleon -> settings.json
  - REMOTE-URL

- Python script testing
  - profiles|extensions create|modify|delete
  - test chameleon rules for proxy|headers
  - test fingerprintjs

## Debugging

- Check Firefox Installed Extension IDs
  - about:debugging#/runtime/this-firefox
  - about:config -> extensions.webextensions.uuids
