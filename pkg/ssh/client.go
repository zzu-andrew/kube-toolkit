package ssh

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
)

// SSHClient 定义一个结构体来保存 SSH 客户端配置
type SSHClient struct {
	Client *ssh.Client
}

// NewSSHClient 创建一个新的 SSH 客户端
func NewSSHClient(user, host, privateKeyPath string) (*SSHClient, error) {
	// 读取私钥文件
	key, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read private key: %v", err)
	}

	// 解析私钥
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %v", err)
	}

	// 配置 SSH 客户端
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // 忽略主机密钥验证（仅用于测试环境）
	}

	// 连接到远程服务器
	client, err := ssh.Dial("tcp", host, config)
	if err != nil {
		return nil, fmt.Errorf("failed to dial: %v", err)
	}

	return &SSHClient{Client: client}, nil
}

// RunCommand 在远程服务器上执行命令并返回输出
func (sc *SSHClient) RunCommand(command string) (string, error) {
	// 创建一个新的会话
	session, err := sc.Client.NewSession()
	if err != nil {
		return "", fmt.Errorf("failed to create session: %v", err)
	}
	defer session.Close()

	// 执行命令并捕获输出
	output, err := session.CombinedOutput(command)
	if err != nil {
		return "", fmt.Errorf("failed to run command: %v", err)
	}

	return string(output), nil
}

func Dat() {
	// 配置 SSH 连接信息
	user := "your-username"                  // 替换为您的用户名
	host := "your-server-ip:22"              // 替换为您的服务器地址和端口
	privateKeyPath := "/path/to/private/key" // 替换为您的私钥路径

	// 创建 SSH 客户端
	client, err := NewSSHClient(user, host, privateKeyPath)
	if err != nil {
		log.Fatalf("Failed to create SSH client: %v", err)
	}
	defer client.Client.Close()

	// 执行 kubectl 命令
	command := "kubectl get pod -A"
	output, err := client.RunCommand(command)
	if err != nil {
		log.Fatalf("Failed to execute command: %v", err)
	}

	// 输出命令结果
	fmt.Println("Command Output:")
	fmt.Println(output)
}
