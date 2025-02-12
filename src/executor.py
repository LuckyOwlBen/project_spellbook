def execute_script(command: str) -> str:
    import subprocess

    result = subprocess.run(command, shell=True, capture_output=True, text=True)
    return result.stdout.strip() + "\n" + result.stderr.strip() if result.stderr else result.stdout.strip()