package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

type serviceAccount struct {
	Name      string
	Namespace string
}

type clusterRoleBinding struct {
	SvcAcc *serviceAccount
	Name   string
}

func main() {
	//configure service account
	svcAcc := serviceAccount{"yes-man", "vault-tec"}
	serviceAccountContent := fmt.Sprintf(`---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: %s
  namespace: %s`, svcAcc.Name, svcAcc.Namespace)

	//configure cluster role binding
	cRlBnd := clusterRoleBinding{&svcAcc, "yes-man"}
	clusterRoleBindingContent := fmt.Sprintf(`---
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: %s
subjects:
- kind: ServiceAccount
  name: %s
  namespace: %s
roleRef:
  kind: ClusterRole
  name: view
  apiGroup: rbac.authorization.k8s.io`, cRlBnd.Name, cRlBnd.SvcAcc.Name, cRlBnd.SvcAcc.Namespace)
	fmt.Println(serviceAccountContent)
	fmt.Println(clusterRoleBindingContent)

	//create File
	serviceAccountFile, err := os.Create("serviceAccount.yaml")
	if err != nil {
		panic(err)
	}

	clusterRoleBindingFile, err := os.Create("clusterRoleBinding.yaml")
	if err != nil {
		panic(err)
	}
	//create byte slices that will be written to the files
	serviceAccountBytes := []byte(serviceAccountContent)
	clusterRoleBindingBytes := []byte(clusterRoleBindingContent)

	//finally write the bite slices to the two files
	serviceAccountWrite, err := serviceAccountFile.Write(serviceAccountBytes)
	if serviceAccountWrite != len(serviceAccountBytes) {
		panic(err)
	}

	clusterRoleBindingWrite, err := clusterRoleBindingFile.Write(clusterRoleBindingBytes)
	if clusterRoleBindingWrite != len(clusterRoleBindingBytes) {
		panic(err)
	}

	//now we invoke kubectl to apply any changes
	applyYAML("serviceAccount.yaml")
	applyYAML("clusterRoleBinding.yaml")
	printToken(svcAcc.Name, svcAcc.Namespace)
}

func applyYAML(YAMLFile string) {
	out, err := exec.Command("kubectl", "apply", "-f", YAMLFile).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))
}

func printToken(serviceName, serviceNamespace string) {
	out, err := exec.Command("kubectl", "-n", serviceNamespace, "describe", "secret", "$(kubectl -n kube-system get secret | grep "+serviceName+" | awk '{print $1}')").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))
}
