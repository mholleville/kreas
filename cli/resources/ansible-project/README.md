# Ansible project by Kreas

The following elements constitute the project :

```
.
├── group_vars/            # contains the variables and vault (encrypted variables)
├── roles/                  # contains an example role following the ansible best practices
├── playbook.yaml           # is the input file when we lauching the ansible command
├── sample                  # is the example inventory
└── README.md               # OPTIONAL: A human-readable README file
```

### Usage

```shell script
$ ansible-playbook -i sample --ask-vault-pass playbook.yaml
```

NOTE : The `vault password` is `sample`

### Usefull links

- [ansible](https://docs.ansible.com)
- [concepts](https://docs.ansible.com/ansible/latest/network/getting_started/basic_concepts.html)
- [ansible vault](https://docs.ansible.com/ansible/latest/user_guide/vault.html)
- [inventory](https://docs.ansible.com/ansible/latest/user_guide/intro_inventory.html)
- [roles](https://docs.ansible.com/ansible/latest/user_guide/playbooks_reuse_roles.html)
- [playbooks](https://docs.ansible.com/ansible/latest/user_guide/playbooks.html)
- [best practices](https://docs.ansible.com/ansible/latest/user_guide/playbooks_best_practices.html)

