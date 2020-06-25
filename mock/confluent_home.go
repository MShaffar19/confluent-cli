// Code generated by mocker. DO NOT EDIT.
// github.com/travisjeffery/mocker
// Source: confluent_home.go

package mock

import (
	sync "sync"
)

// MockConfluentHome is a mock of ConfluentHome interface
type MockConfluentHome struct {
	lockGetFile sync.Mutex
	GetFileFunc func(path ...string) (string, error)

	lockHasFile sync.Mutex
	HasFileFunc func(path ...string) (bool, error)

	lockFindFile sync.Mutex
	FindFileFunc func(pattern string) ([]string, error)

	lockIsConfluentPlatform sync.Mutex
	IsConfluentPlatformFunc func() (bool, error)

	lockGetConfluentVersion sync.Mutex
	GetConfluentVersionFunc func() (string, error)

	lockGetServiceStartScript sync.Mutex
	GetServiceStartScriptFunc func(service string) (string, error)

	lockReadServiceConfig sync.Mutex
	ReadServiceConfigFunc func(service string) ([]byte, error)

	lockReadServicePort sync.Mutex
	ReadServicePortFunc func(service string) (int, error)

	lockGetVersion sync.Mutex
	GetVersionFunc func(service string) (string, error)

	lockGetConnectorConfigFile sync.Mutex
	GetConnectorConfigFileFunc func(connector string) (string, error)

	lockGetKafkaScript sync.Mutex
	GetKafkaScriptFunc func(mode, format string) (string, error)

	lockReadDemoReadme sync.Mutex
	ReadDemoReadmeFunc func(demo string) (string, error)

	calls struct {
		GetFile []struct {
			Path []string
		}
		HasFile []struct {
			Path []string
		}
		FindFile []struct {
			Pattern string
		}
		IsConfluentPlatform []struct {
		}
		GetConfluentVersion []struct {
		}
		GetServiceStartScript []struct {
			Service string
		}
		ReadServiceConfig []struct {
			Service string
		}
		ReadServicePort []struct {
			Service string
		}
		GetVersion []struct {
			Service string
		}
		GetConnectorConfigFile []struct {
			Connector string
		}
		GetKafkaScript []struct {
			Mode   string
			Format string
		}
		ReadDemoReadme []struct {
			Demo string
		}
	}
}

// GetFile mocks base method by wrapping the associated func.
func (m *MockConfluentHome) GetFile(path ...string) (string, error) {
	m.lockGetFile.Lock()
	defer m.lockGetFile.Unlock()

	if m.GetFileFunc == nil {
		panic("mocker: MockConfluentHome.GetFileFunc is nil but MockConfluentHome.GetFile was called.")
	}

	call := struct {
		Path []string
	}{
		Path: path,
	}

	m.calls.GetFile = append(m.calls.GetFile, call)

	return m.GetFileFunc(path...)
}

// GetFileCalled returns true if GetFile was called at least once.
func (m *MockConfluentHome) GetFileCalled() bool {
	m.lockGetFile.Lock()
	defer m.lockGetFile.Unlock()

	return len(m.calls.GetFile) > 0
}

// GetFileCalls returns the calls made to GetFile.
func (m *MockConfluentHome) GetFileCalls() []struct {
	Path []string
} {
	m.lockGetFile.Lock()
	defer m.lockGetFile.Unlock()

	return m.calls.GetFile
}

// HasFile mocks base method by wrapping the associated func.
func (m *MockConfluentHome) HasFile(path ...string) (bool, error) {
	m.lockHasFile.Lock()
	defer m.lockHasFile.Unlock()

	if m.HasFileFunc == nil {
		panic("mocker: MockConfluentHome.HasFileFunc is nil but MockConfluentHome.HasFile was called.")
	}

	call := struct {
		Path []string
	}{
		Path: path,
	}

	m.calls.HasFile = append(m.calls.HasFile, call)

	return m.HasFileFunc(path...)
}

// HasFileCalled returns true if HasFile was called at least once.
func (m *MockConfluentHome) HasFileCalled() bool {
	m.lockHasFile.Lock()
	defer m.lockHasFile.Unlock()

	return len(m.calls.HasFile) > 0
}

// HasFileCalls returns the calls made to HasFile.
func (m *MockConfluentHome) HasFileCalls() []struct {
	Path []string
} {
	m.lockHasFile.Lock()
	defer m.lockHasFile.Unlock()

	return m.calls.HasFile
}

// FindFile mocks base method by wrapping the associated func.
func (m *MockConfluentHome) FindFile(pattern string) ([]string, error) {
	m.lockFindFile.Lock()
	defer m.lockFindFile.Unlock()

	if m.FindFileFunc == nil {
		panic("mocker: MockConfluentHome.FindFileFunc is nil but MockConfluentHome.FindFile was called.")
	}

	call := struct {
		Pattern string
	}{
		Pattern: pattern,
	}

	m.calls.FindFile = append(m.calls.FindFile, call)

	return m.FindFileFunc(pattern)
}

// FindFileCalled returns true if FindFile was called at least once.
func (m *MockConfluentHome) FindFileCalled() bool {
	m.lockFindFile.Lock()
	defer m.lockFindFile.Unlock()

	return len(m.calls.FindFile) > 0
}

// FindFileCalls returns the calls made to FindFile.
func (m *MockConfluentHome) FindFileCalls() []struct {
	Pattern string
} {
	m.lockFindFile.Lock()
	defer m.lockFindFile.Unlock()

	return m.calls.FindFile
}

// IsConfluentPlatform mocks base method by wrapping the associated func.
func (m *MockConfluentHome) IsConfluentPlatform() (bool, error) {
	m.lockIsConfluentPlatform.Lock()
	defer m.lockIsConfluentPlatform.Unlock()

	if m.IsConfluentPlatformFunc == nil {
		panic("mocker: MockConfluentHome.IsConfluentPlatformFunc is nil but MockConfluentHome.IsConfluentPlatform was called.")
	}

	call := struct {
	}{}

	m.calls.IsConfluentPlatform = append(m.calls.IsConfluentPlatform, call)

	return m.IsConfluentPlatformFunc()
}

// IsConfluentPlatformCalled returns true if IsConfluentPlatform was called at least once.
func (m *MockConfluentHome) IsConfluentPlatformCalled() bool {
	m.lockIsConfluentPlatform.Lock()
	defer m.lockIsConfluentPlatform.Unlock()

	return len(m.calls.IsConfluentPlatform) > 0
}

// IsConfluentPlatformCalls returns the calls made to IsConfluentPlatform.
func (m *MockConfluentHome) IsConfluentPlatformCalls() []struct {
} {
	m.lockIsConfluentPlatform.Lock()
	defer m.lockIsConfluentPlatform.Unlock()

	return m.calls.IsConfluentPlatform
}

// GetConfluentVersion mocks base method by wrapping the associated func.
func (m *MockConfluentHome) GetConfluentVersion() (string, error) {
	m.lockGetConfluentVersion.Lock()
	defer m.lockGetConfluentVersion.Unlock()

	if m.GetConfluentVersionFunc == nil {
		panic("mocker: MockConfluentHome.GetConfluentVersionFunc is nil but MockConfluentHome.GetConfluentVersion was called.")
	}

	call := struct {
	}{}

	m.calls.GetConfluentVersion = append(m.calls.GetConfluentVersion, call)

	return m.GetConfluentVersionFunc()
}

// GetConfluentVersionCalled returns true if GetConfluentVersion was called at least once.
func (m *MockConfluentHome) GetConfluentVersionCalled() bool {
	m.lockGetConfluentVersion.Lock()
	defer m.lockGetConfluentVersion.Unlock()

	return len(m.calls.GetConfluentVersion) > 0
}

// GetConfluentVersionCalls returns the calls made to GetConfluentVersion.
func (m *MockConfluentHome) GetConfluentVersionCalls() []struct {
} {
	m.lockGetConfluentVersion.Lock()
	defer m.lockGetConfluentVersion.Unlock()

	return m.calls.GetConfluentVersion
}

// GetServiceStartScript mocks base method by wrapping the associated func.
func (m *MockConfluentHome) GetServiceStartScript(service string) (string, error) {
	m.lockGetServiceStartScript.Lock()
	defer m.lockGetServiceStartScript.Unlock()

	if m.GetServiceStartScriptFunc == nil {
		panic("mocker: MockConfluentHome.GetServiceStartScriptFunc is nil but MockConfluentHome.GetServiceStartScript was called.")
	}

	call := struct {
		Service string
	}{
		Service: service,
	}

	m.calls.GetServiceStartScript = append(m.calls.GetServiceStartScript, call)

	return m.GetServiceStartScriptFunc(service)
}

// GetServiceStartScriptCalled returns true if GetServiceStartScript was called at least once.
func (m *MockConfluentHome) GetServiceStartScriptCalled() bool {
	m.lockGetServiceStartScript.Lock()
	defer m.lockGetServiceStartScript.Unlock()

	return len(m.calls.GetServiceStartScript) > 0
}

// GetServiceStartScriptCalls returns the calls made to GetServiceStartScript.
func (m *MockConfluentHome) GetServiceStartScriptCalls() []struct {
	Service string
} {
	m.lockGetServiceStartScript.Lock()
	defer m.lockGetServiceStartScript.Unlock()

	return m.calls.GetServiceStartScript
}

// ReadServiceConfig mocks base method by wrapping the associated func.
func (m *MockConfluentHome) ReadServiceConfig(service string) ([]byte, error) {
	m.lockReadServiceConfig.Lock()
	defer m.lockReadServiceConfig.Unlock()

	if m.ReadServiceConfigFunc == nil {
		panic("mocker: MockConfluentHome.ReadServiceConfigFunc is nil but MockConfluentHome.ReadServiceConfig was called.")
	}

	call := struct {
		Service string
	}{
		Service: service,
	}

	m.calls.ReadServiceConfig = append(m.calls.ReadServiceConfig, call)

	return m.ReadServiceConfigFunc(service)
}

// ReadServiceConfigCalled returns true if ReadServiceConfig was called at least once.
func (m *MockConfluentHome) ReadServiceConfigCalled() bool {
	m.lockReadServiceConfig.Lock()
	defer m.lockReadServiceConfig.Unlock()

	return len(m.calls.ReadServiceConfig) > 0
}

// ReadServiceConfigCalls returns the calls made to ReadServiceConfig.
func (m *MockConfluentHome) ReadServiceConfigCalls() []struct {
	Service string
} {
	m.lockReadServiceConfig.Lock()
	defer m.lockReadServiceConfig.Unlock()

	return m.calls.ReadServiceConfig
}

// ReadServicePort mocks base method by wrapping the associated func.
func (m *MockConfluentHome) ReadServicePort(service string) (int, error) {
	m.lockReadServicePort.Lock()
	defer m.lockReadServicePort.Unlock()

	if m.ReadServicePortFunc == nil {
		panic("mocker: MockConfluentHome.ReadServicePortFunc is nil but MockConfluentHome.ReadServicePort was called.")
	}

	call := struct {
		Service string
	}{
		Service: service,
	}

	m.calls.ReadServicePort = append(m.calls.ReadServicePort, call)

	return m.ReadServicePortFunc(service)
}

// ReadServicePortCalled returns true if ReadServicePort was called at least once.
func (m *MockConfluentHome) ReadServicePortCalled() bool {
	m.lockReadServicePort.Lock()
	defer m.lockReadServicePort.Unlock()

	return len(m.calls.ReadServicePort) > 0
}

// ReadServicePortCalls returns the calls made to ReadServicePort.
func (m *MockConfluentHome) ReadServicePortCalls() []struct {
	Service string
} {
	m.lockReadServicePort.Lock()
	defer m.lockReadServicePort.Unlock()

	return m.calls.ReadServicePort
}

// GetVersion mocks base method by wrapping the associated func.
func (m *MockConfluentHome) GetVersion(service string) (string, error) {
	m.lockGetVersion.Lock()
	defer m.lockGetVersion.Unlock()

	if m.GetVersionFunc == nil {
		panic("mocker: MockConfluentHome.GetVersionFunc is nil but MockConfluentHome.GetVersion was called.")
	}

	call := struct {
		Service string
	}{
		Service: service,
	}

	m.calls.GetVersion = append(m.calls.GetVersion, call)

	return m.GetVersionFunc(service)
}

// GetVersionCalled returns true if GetVersion was called at least once.
func (m *MockConfluentHome) GetVersionCalled() bool {
	m.lockGetVersion.Lock()
	defer m.lockGetVersion.Unlock()

	return len(m.calls.GetVersion) > 0
}

// GetVersionCalls returns the calls made to GetVersion.
func (m *MockConfluentHome) GetVersionCalls() []struct {
	Service string
} {
	m.lockGetVersion.Lock()
	defer m.lockGetVersion.Unlock()

	return m.calls.GetVersion
}

// GetConnectorConfigFile mocks base method by wrapping the associated func.
func (m *MockConfluentHome) GetConnectorConfigFile(connector string) (string, error) {
	m.lockGetConnectorConfigFile.Lock()
	defer m.lockGetConnectorConfigFile.Unlock()

	if m.GetConnectorConfigFileFunc == nil {
		panic("mocker: MockConfluentHome.GetConnectorConfigFileFunc is nil but MockConfluentHome.GetConnectorConfigFile was called.")
	}

	call := struct {
		Connector string
	}{
		Connector: connector,
	}

	m.calls.GetConnectorConfigFile = append(m.calls.GetConnectorConfigFile, call)

	return m.GetConnectorConfigFileFunc(connector)
}

// GetConnectorConfigFileCalled returns true if GetConnectorConfigFile was called at least once.
func (m *MockConfluentHome) GetConnectorConfigFileCalled() bool {
	m.lockGetConnectorConfigFile.Lock()
	defer m.lockGetConnectorConfigFile.Unlock()

	return len(m.calls.GetConnectorConfigFile) > 0
}

// GetConnectorConfigFileCalls returns the calls made to GetConnectorConfigFile.
func (m *MockConfluentHome) GetConnectorConfigFileCalls() []struct {
	Connector string
} {
	m.lockGetConnectorConfigFile.Lock()
	defer m.lockGetConnectorConfigFile.Unlock()

	return m.calls.GetConnectorConfigFile
}

// GetKafkaScript mocks base method by wrapping the associated func.
func (m *MockConfluentHome) GetKafkaScript(mode, format string) (string, error) {
	m.lockGetKafkaScript.Lock()
	defer m.lockGetKafkaScript.Unlock()

	if m.GetKafkaScriptFunc == nil {
		panic("mocker: MockConfluentHome.GetKafkaScriptFunc is nil but MockConfluentHome.GetKafkaScript was called.")
	}

	call := struct {
		Mode   string
		Format string
	}{
		Mode:   mode,
		Format: format,
	}

	m.calls.GetKafkaScript = append(m.calls.GetKafkaScript, call)

	return m.GetKafkaScriptFunc(mode, format)
}

// GetKafkaScriptCalled returns true if GetKafkaScript was called at least once.
func (m *MockConfluentHome) GetKafkaScriptCalled() bool {
	m.lockGetKafkaScript.Lock()
	defer m.lockGetKafkaScript.Unlock()

	return len(m.calls.GetKafkaScript) > 0
}

// GetKafkaScriptCalls returns the calls made to GetKafkaScript.
func (m *MockConfluentHome) GetKafkaScriptCalls() []struct {
	Mode   string
	Format string
} {
	m.lockGetKafkaScript.Lock()
	defer m.lockGetKafkaScript.Unlock()

	return m.calls.GetKafkaScript
}

// ReadDemoReadme mocks base method by wrapping the associated func.
func (m *MockConfluentHome) ReadDemoReadme(demo string) (string, error) {
	m.lockReadDemoReadme.Lock()
	defer m.lockReadDemoReadme.Unlock()

	if m.ReadDemoReadmeFunc == nil {
		panic("mocker: MockConfluentHome.ReadDemoReadmeFunc is nil but MockConfluentHome.ReadDemoReadme was called.")
	}

	call := struct {
		Demo string
	}{
		Demo: demo,
	}

	m.calls.ReadDemoReadme = append(m.calls.ReadDemoReadme, call)

	return m.ReadDemoReadmeFunc(demo)
}

// ReadDemoReadmeCalled returns true if ReadDemoReadme was called at least once.
func (m *MockConfluentHome) ReadDemoReadmeCalled() bool {
	m.lockReadDemoReadme.Lock()
	defer m.lockReadDemoReadme.Unlock()

	return len(m.calls.ReadDemoReadme) > 0
}

// ReadDemoReadmeCalls returns the calls made to ReadDemoReadme.
func (m *MockConfluentHome) ReadDemoReadmeCalls() []struct {
	Demo string
} {
	m.lockReadDemoReadme.Lock()
	defer m.lockReadDemoReadme.Unlock()

	return m.calls.ReadDemoReadme
}

// Reset resets the calls made to the mocked methods.
func (m *MockConfluentHome) Reset() {
	m.lockGetFile.Lock()
	m.calls.GetFile = nil
	m.lockGetFile.Unlock()
	m.lockHasFile.Lock()
	m.calls.HasFile = nil
	m.lockHasFile.Unlock()
	m.lockFindFile.Lock()
	m.calls.FindFile = nil
	m.lockFindFile.Unlock()
	m.lockIsConfluentPlatform.Lock()
	m.calls.IsConfluentPlatform = nil
	m.lockIsConfluentPlatform.Unlock()
	m.lockGetConfluentVersion.Lock()
	m.calls.GetConfluentVersion = nil
	m.lockGetConfluentVersion.Unlock()
	m.lockGetServiceStartScript.Lock()
	m.calls.GetServiceStartScript = nil
	m.lockGetServiceStartScript.Unlock()
	m.lockReadServiceConfig.Lock()
	m.calls.ReadServiceConfig = nil
	m.lockReadServiceConfig.Unlock()
	m.lockReadServicePort.Lock()
	m.calls.ReadServicePort = nil
	m.lockReadServicePort.Unlock()
	m.lockGetVersion.Lock()
	m.calls.GetVersion = nil
	m.lockGetVersion.Unlock()
	m.lockGetConnectorConfigFile.Lock()
	m.calls.GetConnectorConfigFile = nil
	m.lockGetConnectorConfigFile.Unlock()
	m.lockGetKafkaScript.Lock()
	m.calls.GetKafkaScript = nil
	m.lockGetKafkaScript.Unlock()
	m.lockReadDemoReadme.Lock()
	m.calls.ReadDemoReadme = nil
	m.lockReadDemoReadme.Unlock()
}
