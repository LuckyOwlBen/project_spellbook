from executor import execute_script

if __name__ == "__main__":
    command = "ls -la"
    print(execute_script(command))